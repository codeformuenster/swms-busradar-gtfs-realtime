package busradar

import (
	"fmt"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"github.com/geops/gtfsparser"
	gtfsparser_gtfs "github.com/geops/gtfsparser/gtfs"
)

type Feature struct {
	Geometry struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Visfahrplanlagezst int    `json:"visfahrplanlagezst"`
		Linienid           string `json:"linienid"`
		Nachhst            string `json:"nachhst"`
		Fahrtstatus        string `json:"fahrtstatus"`
		Richtungstext      string `json:"richtungstext"`
		Sequenz            int    `json:"sequenz"`
		Delay              int    `json:"delay"`
		Linientext         string `json:"linientext"`
		Fahrtbezeichner    string `json:"fahrtbezeichner"`
		Richtungsid        string `json:"richtungsid"`
		Fahrzeugid         string `json:"fahrzeugid"`
		Betriebstag        string `json:"betriebstag"`
		Akthst             string `json:"akthst"`
		Operation          string `json:"operation,omitempty"`
	} `json:"properties"`
	tripID string
}

func (f *Feature) TripDescriptor() *gtfs.TripDescriptor {
	routeID := f.Properties.Linientext
	tripID := f.tripID
	// directionID64, _ := strconv.ParseUint(f.Properties.Richtungsid, 10, 32)
	// directionID := uint32(directionID64)
	td := gtfs.TripDescriptor{
		TripId:  &tripID,
		RouteId: &routeID,
		// DirectionId:          &directionID,
		ScheduleRelationship: gtfs.TripDescriptor_SCHEDULED.Enum(),
	}
	return &td
}

func (f *Feature) VehicleDescriptor() *gtfs.VehicleDescriptor {
	idStr := fmt.Sprintf("%s_%s_%s", f.Properties.Fahrzeugid, f.Properties.Richtungsid, f.Properties.Fahrtbezeichner)
	labelStr := fmt.Sprintf("Linie %s Richtung %s", f.Properties.Linientext, f.Properties.Richtungstext)
	v := gtfs.VehicleDescriptor{
		Id:    &idStr,
		Label: &labelStr,
	}
	return &v
}

func (f *Feature) Timestamp() *uint64 {
	ts := uint64(f.Properties.Visfahrplanlagezst)
	return &ts
}

func (f *Feature) Delay() *int32 {
	d := int32(f.Properties.Delay)
	return &d
}

func (f *Feature) StopTimeEvent() *gtfs.TripUpdate_StopTimeEvent {
	stopTimeTime := int64(*f.Timestamp())
	ste := gtfs.TripUpdate_StopTimeEvent{
		Delay: f.Delay(),
		Time:  &stopTimeTime,
	}
	return &ste
}

func (f *Feature) StopSequence() *uint32 {
	stopSequence := uint32(f.Properties.Sequenz)
	return &stopSequence
}

func (f *Feature) StopId() *string {
	stopID := f.Properties.Akthst
	return &stopID
}

func (f *Feature) StopTimeUpdate() []*gtfs.TripUpdate_StopTimeUpdate {
	stus := []*gtfs.TripUpdate_StopTimeUpdate{}

	stu := gtfs.TripUpdate_StopTimeUpdate{
		Arrival: f.StopTimeEvent(),
		StopId:  f.StopId(),
		// StopSequence:         f.StopSequence(),
		ScheduleRelationship: gtfs.Default_TripUpdate_StopTimeUpdate_ScheduleRelationship.Enum(),
	}

	stus = append(stus, &stu)

	return stus
}

func (f *Feature) Position() *gtfs.Position {
	lon := float32(f.Geometry.Coordinates[0])
	lat := float32(f.Geometry.Coordinates[1])
	p := gtfs.Position{
		Latitude:  &lat,
		Longitude: &lon,
	}

	return &p
}

func (f *Feature) Id() *string {
	id := fmt.Sprintf("%s_%s_%s_%d",
		f.Properties.Linienid,
		f.Properties.Fahrzeugid,
		f.Properties.Richtungsid,
		f.Properties.Visfahrplanlagezst,
	)
	return &id
}

// MatchGTFSTrip tries to match values from the Feature to a Trip inside the
// static GTFS feed from a list of Trips
func (f *Feature) MatchGTFSTrip(trips map[string]*gtfsparser_gtfs.Trip) error {
	for tripID, trip := range trips {
		if trip.Route.Id == f.Properties.Linientext && trip.Headsign == f.Properties.Richtungstext {
			f.tripID = tripID
			return nil
		}
	}
	return fmt.Errorf("Couldn't find matching trip for %s and %s",
		f.Properties.Linientext,
		f.Properties.Richtungstext,
	)
}

func (f *Feature) FeedEntity(feed *gtfsparser.Feed, creationTime uint64) (*gtfs.FeedEntity, error) {
	// shared stuff
	if _, ok := feed.Stops[*f.StopId()]; ok == false {
		return &gtfs.FeedEntity{}, fmt.Errorf("Couldn't find matching stop id for %s", *f.StopId())
	}
	err := f.MatchGTFSTrip(feed.Trips)
	if err != nil {
		return &gtfs.FeedEntity{}, err
	}
	tripDescriptor := f.TripDescriptor()

	vehiclePosition := gtfs.VehiclePosition{
		Trip:                tripDescriptor,
		Vehicle:             f.VehicleDescriptor(),
		Timestamp:           &creationTime,
		Position:            f.Position(),
		CurrentStopSequence: f.StopSequence(),
		StopId:              f.StopId(),
	}

	entity := gtfs.FeedEntity{
		Id:      f.Id(),
		Vehicle: &vehiclePosition,
	}

	tripUpdate := gtfs.TripUpdate{
		Trip:           tripDescriptor,
		Vehicle:        f.VehicleDescriptor(),
		Timestamp:      &creationTime,
		Delay:          f.Delay(),
		StopTimeUpdate: f.StopTimeUpdate(),
	}
	entity.TripUpdate = &tripUpdate

	return &entity, nil
}
