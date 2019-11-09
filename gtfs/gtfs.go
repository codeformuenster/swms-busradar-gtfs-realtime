package gtfs

import (
	"fmt"
	"os"

	"github.com/geops/gtfsparser"
)

var feed *gtfsparser.Feed

func init() {
	feedFilepath := "gtfsfeed"
	if os.Getenv("SWMS_GTFS_PATH") != "" {
		feedFilepath = os.Getenv("SWMS_GTFS_PATH")
	}
	fmt.Printf("Initializing GTFS feed from path \"%s\" ... ", feedFilepath)

	feed = gtfsparser.NewFeed()
	err := feed.Parse(feedFilepath)
	if err != nil {
		fmt.Printf("Error parsing %v\n", err)
		return
	}
	fmt.Println("done")
}

func Feed() *gtfsparser.Feed {
	return feed
}
