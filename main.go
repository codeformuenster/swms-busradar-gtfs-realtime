package main

import (
	"encoding/json"
	"fmt"

	"github.com/codeformuenster/swms-busradar-gtfs-realtime/busradar"
	"github.com/codeformuenster/swms-busradar-gtfs-realtime/gtfs"
)

func main() {
	initial, err := busradar.NewResponseFromStatic()
	if err != nil {
		fmt.Println(err)
		return
	}
	feed, err := initial.GTFSRealtimeFeedMessage()
	if err != nil {
		fmt.Println(err)
		return
	}

	// initFeed()
	str, err := json.Marshal(feed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(str))

	gtfs.Feed()
	// fmt.Println(string(a))
}
