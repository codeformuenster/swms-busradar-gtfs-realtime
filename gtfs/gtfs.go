package gtfs

import (
	"github.com/geops/gtfsparser"
)

func IngestGTFSFeed(feedFilepath string) (*gtfsparser.Feed, error) {
	feed := gtfsparser.NewFeed()
	err := feed.Parse(feedFilepath)
	if err != nil {
		return feed, err
	}
	return feed, nil
}
