package police

type Officer struct {
	Bio     string         `json:"bio"`
	Contact ContactDetails `json:"contact_details"`
	Name    string         `json:"name"`
	Rank    string         `json:"rank"`
}

func (c Client) Officers(id string) ([]Officer, error) {
	var officers []Officer
	err := c.decodeJSONResponse(id+"/people", &officers)
	if err != nil {
		return nil, err
	}
	return officers, nil
}
