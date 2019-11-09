package busradar

import (
	"time"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"github.com/asmcos/requests"
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

func (s *Response) GTFSRealtimeFeedMessage() (*gtfs.FeedMessage, error) {
	feed := gtfs.FeedMessage{}

	creationTime := uint64(s.CreationTime.Unix())
	version := "2.0"
	header := gtfs.FeedHeader{
		Timestamp:           &creationTime,
		Incrementality:      gtfs.FeedHeader_FULL_DATASET.Enum(),
		GtfsRealtimeVersion: &version,
	}
	feed.Header = &header

	entities := []*gtfs.FeedEntity{}

	t := true
	id := "go"
	e1 := gtfs.FeedEntity{Id: &id, IsDeleted: &t}
	id2 := "go1"
	e2 := gtfs.FeedEntity{Id: &id2, IsDeleted: &t}
	entities = append(entities, &e1)
	entities = append(entities, &e2)
	feed.Entity = entities

	return &feed, nil
}
