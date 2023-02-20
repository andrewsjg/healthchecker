package healthchecks

import (
	"log"
)

func DoHealthChecks(chkConfig Healthchecks, testmode bool) error {
	log.Println("Doing healthchecks")

	checkSuccess := false
	msg := ""

	var err error

	for _, checkDefs := range chkConfig.Checks {
		for _, checkDef := range checkDefs {

			// Perform the health checks
			for _, check := range checkDef.Checks {
				switch check["type"] {
				case "ping":

					target := check["target"]

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
			}

			// Perform the actions

			for _, action := range checkDef.Actions {
				switch action["type"] {
				case "healthcheck.io":

					if checkSuccess {
						log.Printf("Updating healthcheck.io. Msg: %s\n", msg)

						if action["pingurl"] != "" {
							if !testmode {
								updateHealthCheckIO(action["pingurl"], msg)
							} else {
								log.Printf("TEST MODE - Would have run healthcheck.io ping\n")
							}
						}

					} else {
						log.Printf("Healthcheck FAILED. Updating healthcheck.io. Msg: %s\n", msg)

						if action["pingurl"] != "" {
							if !testmode {
								updateHealthCheckIO(action["pingurl"]+"/fail", msg)
							} else {
								log.Printf("TEST MODE - Would have run healthcheck.io fail ping\n")
							}

						}
					}

				case "test":
					if checkSuccess {
						log.Printf("TEST ACTION: Healthcheck. Msg: %s\n", msg)
					} else {
						log.Printf("TEST ACTION: Healthcheck FAILED. Msg: %s\n", msg)
					}
				}
			}
		}
	}

	return nil
}
