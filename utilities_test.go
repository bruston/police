package police

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func newDummyServer(body []byte, statusCode int) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")
		w.WriteHeader(statusCode)
		w.Write(body)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
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
