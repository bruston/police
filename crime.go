package police

import (
	"bytes"
	"fmt"
	"strings"
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
	ID   string `json:"url"`
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
	err := c.decodeJSONResponse("GET", "crime-categories", &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func allIfEmpty(category string) string {
	if category == "" {
		category = "all-crime"
	}
	return category
}

func (c Client) StreetCrime(lat, long float64, date string, category string) ([]Crime, error) {
	category = allIfEmpty(category)
	url := fmt.Sprintf("crimes-street/%s?lat=%f&lng=%f&date=%s", category, lat, long, date)
	var crimes []Crime
	if err := c.decodeJSONResponse("GET", url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}

type Coordinate struct {
	Latitude, Longitude float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
}

func (c Client) StreetCrimePoly(coordinates []Coordinate, category, date string) ([]Crime, error) {
	category = allIfEmpty(category)
	var buf bytes.Buffer
	for _, v := range coordinates {
		buf.WriteString(v.String() + ":")
	}
	poly := strings.TrimSuffix(buf.String(), ":")
	url := fmt.Sprintf("crimes-street/%s?poly=%s&date=%s", category, poly, date)
	if date == "" {
		url = strings.TrimSuffix(url, "&date=")
	}
	var crimes []Crime
	if err := c.decodeJSONResponse("POST", url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}

func (c Client) CrimesAtLocation(id uint64, date string) ([]Crime, error) {
	url := fmt.Sprintf("crimes-at-location?date=%s&location=%d", date, id)
	var crimes []Crime
	if err := c.decodeJSONResponse("GET", url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}

func (c Client) CrimesClosestLocation(latitude, longitude float64, date string) ([]Crime, error) {
	url := fmt.Sprintf("crimes-at-location?date=%s&lat=%f&lng=%f", date, latitude, longitude)
	var crimes []Crime
	if err := c.decodeJSONResponse("GET", url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}

func (c Client) CrimesNoLocation(category, forceID, date string) ([]Crime, error) {
	category = allIfEmpty(category)
	url := fmt.Sprintf("crimes-no-location?category=%s&force=%s&date=%s", category, forceID, date)
	var crimes []Crime
	if err := c.decodeJSONResponse("GET", url, &crimes); err != nil {
		return nil, err
	}
	return crimes, nil
}
