package api

import (
	"net/http"
	"net/url"
	"testing"
)

func TestParameterParser(t *testing.T) {
	url := &url.URL{
		RawQuery: "q[name_eq]=cone&q[last_name_eq]=Gutierrez",
	}

	request := &http.Request{
		URL: url,
	}

	parser := new(ParameterParser)

	err := parser.Parse(request)

	if err != nil {
		t.Error("An error has ocurred:", err)
	}

	expected := "name = 'cone' AND last_name = 'Gutierrez'"

	if parser.gransakQuery != expected {
		t.Errorf("Mismatch, got: %s expected: %s", parser.gransakQuery, expected)
	}

}
