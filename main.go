package main

import (
	"fmt"
	"io/ioutil"

	"github.com/codeformuenster/swms-busradar-gtfs-realtime/busradar"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	staticResponse, err := busradar.NewResponseFromStatic()
	if err != nil {
		fmt.Println(err)
		return
	}

	initialRealtimeFeed, err := staticResponse.GTFSRealtimeFeedMessage()
	if err != nil {
		fmt.Println(err)
		return
	}

	pb, err := proto.Marshal(initialRealtimeFeed)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("feed", pb, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
