package police

import ()

type Forces struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Force struct {
	Description       string             `json:"description"`
	URL               string             `json:"url"`
	EngagementMethods []EngagementMethod `json:"engagement_methods"`
	Telephone         string             `json:"telephone"`
	ID                string             `json:"id"`
	Name              string             `json:"name"`
}

type EngagementMethod struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

func (f Force) Engagement() []EngagementMethod {
	return f.EngagementMethods
}

func (c Client) Forces() ([]Forces, error) {
	var forces []Forces
	err := c.decodeJSONResponse("forces", &forces)
	if err != nil {
		return nil, err
	}
	return forces, nil
}

func (c Client) Force(id string) (Force, error) {
	var force Force
	err := c.decodeJSONResponse("forces/"+id, &force)
	if err != nil {
		return Force{}, err
	}
	return force, nil
}
