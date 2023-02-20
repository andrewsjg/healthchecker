package hcweb

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/spf13/viper"

	"github.com/andrewsjg/healthchecker/healthchecks"
)

//go:embed hffrontend/dist

var frontend embed.FS

func StartAPI(chkConfig healthchecks.CheckConfig, testmode bool) {
	log.Println("Starting API")

	stripped, err := fs.Sub(frontend, "hffrontend/dist")

	if err != nil {
		log.Fatalln(err)
	}

	frontendFS := http.FileServer(http.FS(stripped))
	http.Handle("/", frontendFS)

	http.Handle("/api/v1/getConfig", http.HandlerFunc(getConfig(chkConfig)))
	http.Handle("/api/v1/setConfig", http.HandlerFunc(setConfig))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", 8474), nil))

}

func getConfig(chkConfig healthchecks.CheckConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//TODO: REFACTOR THIS

		if err := viper.ReadInConfig(); err == nil {
			//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		} else {
			http.Error(w, "could not read config", http.StatusInternalServerError)
		}

		var checkCfg healthchecks.CheckConfig
		err := viper.Unmarshal(&checkCfg)

		if err != nil {
			http.Error(w, "could not read config", http.StatusInternalServerError)

		}

		//j, err := json.Marshal(chkConfig)
		j, err := json.Marshal(checkCfg)

		if err != nil {
			http.Error(w, "could not read config", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, bytes.NewReader(j))

	}
}

func setConfig(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	var cfg map[string]healthchecks.CheckBlock

	if err != nil {
		log.Println("io panic")
		panic(err)
	}

	err = json.Unmarshal(body, &cfg)

	if err != nil {
		log.Println("Panic!")
		panic(err)
	}

	// Testing writing config to a file

	var checkCfg healthchecks.CheckConfig
	err = viper.Unmarshal(&checkCfg)

	if err != nil {
		log.Println("unable to unmarshal config")
	} else {

		log.Println(checkCfg)

		// find the check to update in the config
		chkIdx := -1
		chkKey := ""

		for idx, check := range checkCfg.Checks {
			for key, _ := range check {
				_, ok := cfg[key]

				if ok {
					chkKey = key
					chkIdx = idx
				}
			}
		}

		// found the check to update
		if chkIdx > -1 {

			newChkDef := healthchecks.CheckDef{}
			newCheck := make(map[string]healthchecks.CheckDef)
			newActionBlock := make(map[string]string)
			newCheckBlock := make(map[string]string)

			newChkDef.Description = cfg[chkKey].Description
			newChkDef.Name = cfg[chkKey].Name

			// TODO: The logic for the action and check blocks will need to change
			// when we want to support more than one action block per check

			newCheckBlock["target"] = cfg[chkKey].Check.Target
			newCheckBlock["type"] = cfg[chkKey].Check.Type

			newChkDef.Check = newCheckBlock

			newActionBlock["type"] = cfg[chkKey].Action.Type
			newActionBlock["pingUrl"] = cfg[chkKey].Action.Pingurl

			newChkDef.Action = newActionBlock

			newChkDef.Enabled = cfg[chkKey].Enabled

			newCheck[chkKey] = newChkDef

			checkCfg.Checks[chkIdx] = newCheck

			//currentConfigFilePath := viper.ConfigFileUsed()

			// Backup the config file
			//err = os.Rename(currentConfigFilePath, currentConfigFilePath+".bak")

			if err != nil {
				log.Fatalf("unable to backup the config file")
			}

			// write new config
			//file, err := os.OpenFile(currentConfigFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
			file, err := os.OpenFile("test.yml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				log.Fatalf("error opening/creating file: %v", err)
			}
			defer file.Close()

			enc := yaml.NewEncoder(file)
			err = enc.Encode(checkCfg)

			if err != nil {
				log.Fatalf("error encoding: %v", err)
			}
		}
	}

}
