package police

import (
	"fmt"
)

type Crime struct {
	Category        string        `json:"category"`
	PersistentID    string        `json:"persistent_id"`
	Month           string        `json:"month"`
	Location        Location      `json:"location"`
	Context         string        `json:"context"`
	ID              uint64        `json:"id"`
	LocationType    string        `json:"location_type"`
	LocationSubtype string        `json:"location_subtype"`
	Outcome         OutcomeStatus `json:"outcome_status"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Street    `json:"street"`
}

type Street struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type OutcomeStatus struct {
	Category string `json:"category"`
	Date     string `json:"date"`
}

func (c Client) StreetCrime(lat, long string, date string) ([]Crime, error) {
	url := fmt.Sprintf("crimes-street/all-crime?lat=%s&lng=%s&date=%s", lat, long, date)
	var crimes []Crime
	if err := c.decodeJSONResponse(url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}
