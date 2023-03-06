package healthchecks

import (
	"errors"
	"fmt"
	"net/http"
)

func httpCheck(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s returned %d", url, resp.StatusCode))
	}

	return nil
}
