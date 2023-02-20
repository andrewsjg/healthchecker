/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/andrewsjg/healthchecker/hcweb"
	"github.com/andrewsjg/healthchecker/healthchecks"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var serverMode bool
var testMode bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "healthchecker",
	Short: "Runs healthchecks against targets defined in a simple config",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Should probably just use the global viper instance
		// instead of passing around the config object
		var checkCfg healthchecks.Healthchecks
		err := viper.Unmarshal(&checkCfg)

		if err != nil {
			log.Fatal("Unable to parse config: ", err)
		}

		err = healthchecks.DoHealthChecks(checkCfg, testMode)
		cobra.CheckErr(err)

		if serverMode {
			hcweb.StartAPI(checkCfg, testMode)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.healthchecker/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&serverMode, "server", "s", false, "Help message for server")

	// Dont run the checks, just report which checks will run
	rootCmd.Flags().BoolVarP(&testMode, "test", "t", false, "Help message for test")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".healthchecker" (without extension).
		viper.AddConfigPath(home + "/.healthchecker")
		viper.AddConfigPath("./.healthchecker")

		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)

		if err := viper.ReadInConfig(); err == nil {
			//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

		} else {
			log.Println("Config error: " + err.Error())
		}
	})

	// If a config file is found, read it in.
	// TODO: Tidy this up
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println("Config error: " + err.Error())
	}
}
