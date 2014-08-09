package police

import ()

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
	Latitutde string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Locations struct {
	Name        string `json:"name"`
	Longitude   string `json:"longitude"`
	Postcode    string `json:"postcode"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (c Client) Neighbourhoods(location string) ([]Neighbourhoods, error) {
	var neighbourhoods []Neighbourhoods
	err := c.decodeJSONResponse(location+"/neighbourhoods", &neighbourhoods)
	if err != nil {
		return nil, err
	}
	return neighbourhoods, nil
}

func (c Client) Neighbourhood(location, id string) (Neighbourhood, error) {
	return Neighbourhood{}, nil
}

func (n Neighbourhoods) Information() (Neighbourhood, error) {
	return Neighbourhood{}, nil
}
