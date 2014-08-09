package police

import ()

type Person struct {
	Bio     string         `json:"bio"`
	Contact ContactDetails `json:"contact_details"`
	Name    string         `json:"name"`
	Rank    string         `json:"rank"`
}

func (c Client) People(id string) ([]Person, error) {
	var people []Person
	err := c.decodeJSONResponse(id+"/people", &people)
	if err != nil {
		return nil, err
	}
	return people, nil
}
