package police

import (
	"reflect"
	"testing"
)

var crimeBody = []byte(`[
{
        "category": "other-crime",
        "context": "",
        "id": 20599348,
        "location": {
            "latitude": "52.634279",
            "longitude": "-1.130922",
            "street": {
                "id": 883366,
                "name": "On or near Granby Place"
            }
        },
        "location_subtype": "",
        "location_type": "Force",
        "month": "2013-01",
        "outcome_status": {
            "category": "Investigation complete; no suspect identified",
            "date": "2013-02"
        },
        "persistent_id": "d9e7fcb89308479a089ad60964ea533a395f54dcff621b04ce08159010b2cd01"
    },
    {
        "category": "other-crime",
        "context": "",
        "id": 20607959,
        "location": {
            "latitude": "52.637354",
            "longitude": "-1.138383",
            "street": {
                "id": 883414,
                "name": "On or near Shopping Area"
            }
        },
        "location_subtype": "",
        "location_type": "Force",
        "month": "2013-01",
        "outcome_status": {
            "category": "Under investigation",
            "date": "2013-01"
        },
        "persistent_id": "cc55bbd38a1cecf1a775569d856d552404a9482b2a48c947e234bd9716224c6c"
    },
    {
        "category": "other-crime",
        "context": "",
        "id": 20597909,
        "location": {
            "latitude": "52.618191",
            "longitude": "-1.136425",
            "street": {
                "id": 882330,
                "name": "On or near Putney Road West"
            }
        },
        "location_subtype": "",
        "location_type": "Force",
        "month": "2013-01",
        "outcome_status": {
            "category": "Investigation complete; no suspect identified",
            "date": "2013-01"
        },
        "persistent_id": "891c558265735c6e8a396d9e6a22ba641b26cff9201d24da322b5c487f22a08b"
    }
]`)

func TestCrimesCoord(t *testing.T) {
	server := newDummyServer(crimeBody, 200)
	defer server.Close()
	p := Client{baseURL: server.URL + "/"}
	crimes, err := p.StreetCrime(100.333, -100.344, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(crimes) != 3 {
		t.Errorf("Expecting slice to be of length 3, it was %d", len(crimes))
	}
	expected := Crime{
		Category: "other-crime",
		ID:       20597909,
		Location: Location{
			Latitude:  52.618191,
			Longitude: -1.136425,
			Street: Street{
				ID:   882330,
				Name: "On or near Putney Road West"},
		},
		LocationType: "Force",
		Month:        "2013-01",
		Outcome: OutcomeStatus{
			Category: "Investigation complete; no suspect identified",
			Date:     "2013-01"},
		PersistentID: "891c558265735c6e8a396d9e6a22ba641b26cff9201d24da322b5c487f22a08b",
	}
	if !reflect.DeepEqual(crimes[2], expected) {
		t.Errorf("expecting %#v, got %#v instead", expected, crimes[2])
	}
}
