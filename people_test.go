package police

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var peopleBody = []byte(`[
    {
        "bio": "A test bio",
        "contact_details": {
            "twitter": "http://www.twitter.com/ACCCLeicsPolice"
        },
        "name": "Joe Bloggs",
        "rank": "Assistant Chief Officer (Crime)"
    }
]`)

func TestPeople(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")
		w.Write(peopleBody)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	p := Client{baseURL: server.URL + "/"}
	people, err := p.People("leicestershire")
	if err != nil {
		t.Fatal(err)
	}
	expected := Person{Bio: "A test bio", Contact: ContactDetails{Twitter: "http://www.twitter.com/ACCCLeicsPolice"}, Name: "Joe Bloggs", Rank: "Assistant Chief Officer (Crime)"}
	if !reflect.DeepEqual(people[0], expected) {
		t.Errorf("expecting %v, got %v instead", expected, people[0])
	}
}
