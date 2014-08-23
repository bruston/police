package police

import (
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
	server := newDummyServer(peopleBody, 200)
	p := Client{baseURL: server.URL + "/"}
	officers, err := p.Officers("leicestershire")
	if err != nil {
		t.Fatal(err)
	}
	expected := Officer{Bio: "A test bio", Contact: ContactDetails{Twitter: "http://www.twitter.com/ACCCLeicsPolice"}, Name: "Joe Bloggs", Rank: "Assistant Chief Officer (Crime)"}
	if !reflect.DeepEqual(officers[0], expected) {
		t.Errorf("expecting %#v, got %#v instead", expected, officers[0])
	}
}
