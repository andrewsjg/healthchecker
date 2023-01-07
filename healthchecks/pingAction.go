package healthchecks

import (
	"errors"
	"fmt"
	"time"

	probing "github.com/andrewsjg/pro-bing"
)

func doPing(target string) (string, error) {

	msg := fmt.Sprintf("%s is down", target)

	pinger, err := probing.NewPinger(target)

	if err != nil {

		return msg + " Reason: " + err.Error(), err
	}

	pinger.Count = 3
	pinger.Timeout = time.Second * 1

	err = pinger.Run() // Blocks until finished.

	if err != nil {
		return msg + " Reason: " + err.Error(), err
	}

	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	//fmt.Printf("Stats for %s\nPackets sent: %d\nPackets recieved: %d\n\n", target, stats.PacketsSent, stats.PacketsRecv)

	if stats.PacketsRecv > 0 {
		msg = fmt.Sprintf("%s is up", target)
	} else {
		err = errors.New("ping loss")
		msg = msg + " Reason: " + err.Error()
	}

	return msg, err
}
