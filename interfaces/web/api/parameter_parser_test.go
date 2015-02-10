package api

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestParameterParser(t *testing.T) {
	url := &url.URL{
		RawQuery: "q[name_eq]=cone&q[last_name_eq]=Gutierrez",
	}

	request := &http.Request{
		URL: url,
	}

	params := NewRequestParameters(request)

	query, err := params.GetGransakQuery()
	if err != nil {
		t.Error("An error has ocurred:", err)
	}

	expected := "name = 'cone'"

	if !strings.Contains(query, expected) {
		t.Errorf("Mismatch, string: %s does not contain: %s", query, expected)
	}

	expected = "last_name = 'Gutierrez'"

	if !strings.Contains(query, expected) {
		t.Errorf("Mismatch, string: %s does not contain: %s", query, expected)
	}
}
