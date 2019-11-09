package busradar

import "github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"

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

func (f *Feature) FeedEntity() (*gtfs.FeedEntity, error) {
	id := "go"
	entity := gtfs.FeedEntity{Id: &id}

	return &entity, nil
}
