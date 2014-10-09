# Police

Police is a Go package that interacts with the REST API for http://data.police.uk which is documented [here](http://data.police.uk/docs/).

## Requirements
- Go 1.3

## Installation

- To install. run `go get github.com/bruston/police`

## Getting Started

The API does not require authentication or an API key, so there's no need to do any configuration. Import the police package and use the New() function to get started.

```Go
package main

import (
	"fmt"
	"github.com/bruston/police"
)

func main() {
	p := police.New()
}
```

## Getting A List Of Police Forces

```Go
forces, err := p.Forces()
if err != nil {
	// Handle errors
}
for _, force := range forces {
	fmt.Printf("ID: %d - Name: %s.\n", force.ID, force.Name)
}
```

## Getting A Specific Police Force

```Go
force, err := p.Force("leicestershire")
if err != nil {
	// Handle errors
}
fmt.Printf("ID: %s - Name: %s - URL: %s - Telephone: %s.\n", force.ID, force.Name, force.URL, force.Telephone)
```

## Getting Force Senior Officer Information

```Go
officers, err := p.Officers("leicestershire")
if err != nil {
        // Handle errors
}
for _, officer := range officers {
        fmt.Printf("Name: %s - Rank: %s.\n", officer.Name, officer.Rank)
}
```

## Getting Street Level Crimes For A Specific Point

To retrieve a list of street level crimes within a one mile radius of a specific point call the `StreetCrime` method.

**Arguments:**

- latitude float64
- longitude float64
- date string (YYYY-MM format)
- category string

If no date is specified, the latest data is returned. If no category ID is specified, all categories are returned.

```Go
crimes, err := p.StreetCrime(52.629729, -1.131592, "2013-01", "")
if err != nil {
	// Handle errors
}
for _, crime := range crimes {
	fmt.Printf("Category: %s - ID: %d - Street: %s - Outcome Status: %s\n", crime.Category, crime.ID, crime.Location.Street.Name, crime.Outcome.Category)
}
```
