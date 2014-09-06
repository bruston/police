package police

type Neighbourhoods struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Neighbourhood struct {
	ForceURL       string `json:"url_force"`
	BoundaryURL    string `json:"url_boundary"`
	ContactDetails `json:"contact_details"`
	Name           string `json:"name"`
	WelcomeMessage string `json:"welcome_message"`
	Links          []Link
	Center         `json:"center"`
	Locations      []Location `json:"locations"`
}

type Link struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

type Center struct {
	Latitutde float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

type Locations struct {
	Name        string  `json:"name"`
	Longitude   float64 `json:"longitude,string"`
	Latitude    float64 `json:"latitude,string"`
	Postcode    string  `json:"postcode"`
	Address     string  `json:"address"`
	Telephone   string  `json:"telephone"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
}

func (c Client) Neighbourhoods(location string) ([]Neighbourhoods, error) {
	var neighbourhoods []Neighbourhoods
	err := c.decodeJSONResponse("GET", location+"/neighbourhoods", &neighbourhoods)
	if err != nil {
		return nil, err
	}
	return neighbourhoods, nil
}

func (c Client) Neighbourhood(location, id string) (Neighbourhood, error) {
	var neighbourhood Neighbourhood
	err := c.decodeJSONResponse("GET", location+"/"+id, &neighbourhood)
	if err != nil {
		return Neighbourhood{}, err
	}
	return neighbourhood, nil
}
