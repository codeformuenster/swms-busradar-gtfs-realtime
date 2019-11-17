package busradar

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"github.com/asmcos/requests"
	"github.com/geops/gtfsparser"
	"github.com/golang/protobuf/proto"
)

type Response struct {
	Features     []Feature `json:"features"`
	CreationTime time.Time
}

func NewResponseFromStatic() (*Response, error) {
	now := time.Now()

	resp, err := requests.Get("https://rest.busradar.conterra.de/prod/fahrzeuge")
	if err != nil {
		return &Response{}, err
	}

	var response Response
	err = resp.Json(&response)
	if err != nil {
		return &Response{}, err
	}

	response.CreationTime = now

	return &response, nil
}

func (s *Response) FeedMessage(feed *gtfsparser.Feed) (*gtfs.FeedMessage, error) {
	feedMessage := gtfs.FeedMessage{}

	creationTime := uint64(s.CreationTime.Unix())
	version := "2.0"
	header := gtfs.FeedHeader{
		Timestamp:           &creationTime,
		Incrementality:      gtfs.FeedHeader_FULL_DATASET.Enum(),
		GtfsRealtimeVersion: &version,
	}
	feedMessage.Header = &header

	entities := []*gtfs.FeedEntity{}

	for _, feature := range s.Features {
		entity, err := feature.FeedEntity(feed, creationTime)
		if err == nil {
			entities = append(entities, entity)
		} else {
			log.Printf("Skipping FeedEntity in FeedMessage: %v\n", err)
		}
	}

	feedMessage.Entity = entities

	return &feedMessage, nil
}

func (s *Response) PersistFeedMessage(staticFeed *gtfsparser.Feed, realtimeFeedPath string) error {
	feed, err := s.FeedMessage(staticFeed)
	if err != nil {
		return err
	}

	pb, err := proto.Marshal(feed)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(realtimeFeedPath, pb, 0644)
	if err != nil {
		return err
	}

	pb, err = json.MarshalIndent(feed, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(realtimeFeedPath+".json", pb, 0644)
	if err != nil {
		return err
	}

	return nil
}
