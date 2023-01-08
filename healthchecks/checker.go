package healthchecks

import "fmt"

func DoHealthChecks(chkConfig CheckConfig) error {

	checkSuccess := false
	msg := ""

	var err error

	for _, checkDefs := range chkConfig.Checks {
		for _, checkDef := range checkDefs {

			switch checkDef.Check["type"] {
			case "ping":
				target := checkDef.Check["target"]
				msg, err = doPing(target)

				if err == nil {
					checkSuccess = true
				} else {
					checkSuccess = false
				}
			}

			switch checkDef.Action["type"] {
			case "healthcheck.io":

				if checkSuccess {
					fmt.Printf("Will send success to healthcheck.io with msg: %s\n", msg)

					if checkDef.Action["pingurl"] != "" {
						updateHealthCheckIO(checkDef.Action["pingurl"], msg)
					}

				} else {
					fmt.Printf("Will send fail to healthcheck.io with msg: %s\n", msg)

					if checkDef.Action["pingurl"] != "" {
						updateHealthCheckIO(checkDef.Action["pingurl"]+"/fail", msg)
					}
				}

			case "test":
				if checkSuccess {
					fmt.Printf("TEST ACTION: The check action succeeded: %s\n", msg)
				} else {
					fmt.Printf("TEST ACTION: The check action failed: %s\n", msg)
				}
			}
		}
	}

	return nil
}
