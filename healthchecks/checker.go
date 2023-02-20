package healthchecks

import (
	"log"
)

func DoHealthChecks(chkConfig CheckConfig, testmode bool) error {

	checkSuccess := false
	msg := ""

	var err error

	for _, checkDefs := range chkConfig.Checks {
		for _, checkDef := range checkDefs {

			switch checkDef.Check["type"] {
			case "ping":

				target := checkDef.Check["target"]

				if !testmode {
					msg, err = doPing(target)
				} else {
					log.Println("TEST MODE - Would have pinged target: " + target)
					err = nil
				}

				if err == nil {
					checkSuccess = true
				} else {
					checkSuccess = false
				}
			}

			switch checkDef.Action["type"] {
			case "healthcheck.io":

				if checkSuccess {
					log.Printf("Healthcheck for %s suceeded. Updating healthcheck.io. Msg: %s\n", checkDef.Check["target"], msg)

					if checkDef.Action["pingurl"] != "" {
						if !testmode {
							updateHealthCheckIO(checkDef.Action["pingurl"], msg)
						} else {
							log.Printf("TEST MODE - Would have run healthcheck.io ping\n")
						}
					}

				} else {
					log.Printf("Healthcheck for %s FAILED. Updating healthcheck.io. Msg: %s\n", checkDef.Check["target"], msg)

					if checkDef.Action["pingurl"] != "" {
						if !testmode {
							updateHealthCheckIO(checkDef.Action["pingurl"]+"/fail", msg)
						} else {
							log.Printf("TEST MODE - Would have run healthcheck.io fail ping\n")
						}

					}
				}

			case "test":
				if checkSuccess {
					log.Printf("TEST ACTION: Healthcheck for %s succeeded. Msg: %s\n", checkDef.Check["target"], msg)
				} else {
					log.Printf("TEST ACTION: Healthcheck for %s FAILED. Msg: %s\n", checkDef.Check["target"], msg)
				}
			}
		}
	}

	return nil
}
