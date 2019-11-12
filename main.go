package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/codeformuenster/swms-busradar-gtfs-realtime/busradar"
	"github.com/codeformuenster/swms-busradar-gtfs-realtime/gtfs"
	"github.com/geops/gtfsparser"
)

func retrieveAndPersistResponse(staticFeed *gtfsparser.Feed, realtimeFeedPath string) error {
	staticResponse, err := busradar.NewResponseFromStatic()
	if err != nil {
		return err
	}

	err = staticResponse.PersistFeedMessage(staticFeed, realtimeFeedPath)
	if err != nil {
		return err
	}

	log.Printf("Persisted realtime feed (%s)\n", realtimeFeedPath)
	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	var feedPath string
	if os.Getenv("GTFS_FEED_PATH") != "" {
		feedPath = os.Getenv("GTFS_FEED_PATH")
	} else {
		cwd, _ := os.Getwd()
		feedPath = path.Join(cwd, "gtfsfeed")
	}
	staticFeed, err := gtfs.IngestGTFSFeed(feedPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Imported static feed (%s)\n", feedPath)

	var realtimeFeedPath string
	if os.Getenv("GTFS_REALTIME_FEED_PATH") != "" {
		realtimeFeedPath = os.Getenv("GTFS_REALTIME_FEED_PATH")
	} else {
		cwd, _ := os.Getwd()
		realtimeFeedPath = path.Join(cwd, "feed")
	}

	err = retrieveAndPersistResponse(staticFeed, realtimeFeedPath)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan struct{})

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := retrieveAndPersistResponse(staticFeed, realtimeFeedPath)

			if err != nil {
				fmt.Println(err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			return
		}
	}
}
