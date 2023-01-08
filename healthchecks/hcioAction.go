package healthchecks

import (
	"fmt"
	"net/http"
	"time"
)

func updateHealthCheckIO(pingURL string, msg string) error {
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	_, err := client.Head(pingURL)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return nil
}
