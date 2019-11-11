package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/codeformuenster/swms-busradar-gtfs-realtime/busradar"
	"github.com/codeformuenster/swms-busradar-gtfs-realtime/gtfs"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	var feedPath string
	if os.Getenv("GTFS_FEED_PATH") != "" {
		feedPath = os.Getenv("GTFS_FEED_PATH")
	} else {
		cwd, _ := os.Getwd()
		feedPath = path.Join(cwd, "gtfsfeed")
	}
	fmt.Printf("Importing static GTFS feed from path %s ... ", feedPath)
	staticFeed, err := gtfs.IngestGTFSFeed(feedPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")

	staticResponse, err := busradar.NewResponseFromStatic()
	if err != nil {
		fmt.Println(err)
		return
	}

	initialRealtimeFeed, err := staticResponse.FeedMessage(staticFeed)
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
