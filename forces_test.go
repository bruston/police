package police

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var forcesBody = []byte(`[
    {
        "id": "avon-and-somerset",
        "name": "Avon and Somerset Constabulary"
    },
    {
        "id": "bedfordshire",
        "name": "Bedfordshire Police"
    },
    {
        "id": "cambridgeshire",
        "name": "Cambridgeshire Constabulary"
    }
]`)

func TestGetForces(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")
		w.Write(forcesBody)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	p := Client{baseURL: server.URL + "/", Client: http.Client{}, UserAgent: USER_AGENT}
	forces, err := p.Forces()
	if err != nil {
		t.Error(err)
	}
	if len(forces) != 3 {
		t.Errorf("Expecting 3 results, got %d.", len(forces))
	}
}

var forceBody = []byte(`{
    "description": "A test description",
    "url": "http://www.leics.police.uk/",
    "engagement_methods": [
        {
            "url": "http://www.facebook.com/pages/Leicester/Leicestershire-Police/76807881169",
            "description": "Become friends with Leicestershire Constabulary",
            "title": "Facebook"
        },
        {
            "url": "http://www.twitter.com/leicspolice",
            "description": "Keep up to date with Leicestershire Constabulary on Twitter",
            "title": "Twitter"
        },
        {
            "url": "http://www.youtube.com/leicspolice",
            "description": "See Leicestershire Constabulary's latest videos on YouTube",
            "title": "YouTube"
        },
        {
            "url": "http://www.leics.police.uk/rss/",
            "description": "Keep informed with Leicestershire Constabulary's RSS feed",
            "title": "RSS"
        }
    ],
    "telephone": "0116 222 2222",
    "id": "leicestershire",
    "name": "Leicestershire Constabulary"
}`)

func TestGetForce(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")
		w.Write(forceBody)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	p := Client{baseURL: server.URL + "/", Client: http.Client{}, UserAgent: USER_AGENT}
	force, err := p.Force("leicestershire")
	if err != nil {
		t.Error(err)
	}
	if force.Name != "Leicestershire Constabulary" {
		t.Errorf("Unexpected name: %s", force.Name)
	}
	if force.Description != "A test description" {
		t.Errorf("Unexpected description: %s", force.Description)
	}
	if len(force.EngagementMethods) != 4 {
		t.Errorf("Expecting 4 EngageMentMethod entries, got %d instead", len(force.EngagementMethods))
	}
	if force.Telephone != "0116 222 2222" {
		t.Errorf("Unexpected telephone number: %s", force.Telephone)
	}
	if force.URL != "http://www.leics.police.uk/" {
		t.Errorf("Unexpected url: %s", force.URL)
	}
}
