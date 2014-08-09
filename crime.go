package police

import (
	"fmt"
	"strconv"
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

type CrimeCategory struct {
	Slug string `json:"url"`
	Name string `json:"name"`
}

type Location struct {
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
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

func (c Client) Categories() ([]CrimeCategory, error) {
	var categories []CrimeCategory
	err := c.decodeJSONResponse("crime-categories", &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c Client) StreetCrime(lat, long float64, date string, category string) ([]Crime, error) {
	if category == "" {
		category = "all-crime"
	}

	url := fmt.Sprintf("crimes-street/%s?lat=%s&lng=%s&date=%s", category, strconv.FormatFloat(lat, 'f', 6, 64), strconv.FormatFloat(long, 'f', 6, 64), date)
	var crimes []Crime
	if err := c.decodeJSONResponse(url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}
