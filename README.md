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
for _, forces := range forces {
	fmt.Printf("ID: %d - Name: %s.\n", forces.ID, forces.Name)
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
