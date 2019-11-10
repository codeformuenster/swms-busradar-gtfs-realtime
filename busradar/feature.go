package busradar

import (
	"fmt"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
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
}

func (f *Feature) TripDescriptor() *gtfs.TripDescriptor {
	td := gtfs.TripDescriptor{
		RouteId: &f.Properties.Linientext,
	}
	return &td
}

func (f *Feature) Vehicle() *gtfs.VehicleDescriptor {
	v := gtfs.VehicleDescriptor{
		Id:    &f.Properties.Fahrzeugid,
		Label: &f.Properties.Richtungstext,
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
		Arrival:      f.StopTimeEvent(),
		StopId:       f.StopId(),
		StopSequence: f.StopSequence(),
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

func (f *Feature) FeedEntity() (*gtfs.FeedEntity, error) {
	// shared stuff
	tripDescriptor := f.TripDescriptor()
	vehicle := f.Vehicle()
	timestamp := f.Timestamp()

	vehiclePosition := gtfs.VehiclePosition{
		Trip:                tripDescriptor,
		Vehicle:             vehicle,
		Timestamp:           timestamp,
		Position:            f.Position(),
		CurrentStopSequence: f.StopSequence(),
		StopId:              f.StopId(),
	}

	entity := gtfs.FeedEntity{
		Id:      f.Id(),
		Vehicle: &vehiclePosition,
	}

	if *f.Delay() != int32(0) {
		tripUpdate := gtfs.TripUpdate{
			Trip:           tripDescriptor,
			Vehicle:        vehicle,
			Timestamp:      timestamp,
			Delay:          f.Delay(),
			StopTimeUpdate: f.StopTimeUpdate(),
		}
		entity.TripUpdate = &tripUpdate
	}

	return &entity, nil
}
