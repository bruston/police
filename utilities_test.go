package police

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type dummyAPI struct {
	statusCode  int
	body        []byte
	lastRequest *http.Request
	*httptest.Server
}

func (d *dummyAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(d.statusCode)
	w.Write(d.body)
	d.lastRequest = r
}

func newDummyServer(body []byte, statusCode int) *dummyAPI {
	s := dummyAPI{statusCode: statusCode, body: body}
	s.Server = httptest.NewServer(&s)
	return &s
}

func newTestClient(baseURL string) Client {
	return Client{BaseURL: baseURL + "/", Client: http.DefaultClient}
}

func TestStructToMap(t *testing.T) {
	var s struct {
		Field1        string
		Field2        string
		IgnoredField  bool
		IgnoredField2 int
		Field3        string
	}
	s.Field1 = "apple"
	s.Field2 = "orange"
	s.Field3 = "pear"
	s.IgnoredField = true
	expected := map[string]string{"field1": "apple", "field2": "orange", "field3": "pear"}
	result := structToMap(s)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expecting %#v, got %#v", expected, result)
	}
}
