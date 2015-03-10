package api

import (
	"net/http"
	"testing"
)

func TestESQueryGenerator_AddToParams(t *testing.T) {
	qg := NewESQueryGenerator("")

	qg.addToParams("name", []string{"foo"})

	if len(qg.params) < 1 {
		t.Error("No params")
		return
	}

	if qg.params[0] != "name:foo" {
		t.Error("Incorrect param value")
		return
	}
}

func TestESQueryGenerator_Parse(t *testing.T) {
	apiUrl := "http://localhost:1440"
	req, err := http.NewRequest("GET", apiUrl+"?field1=1&field2=2", nil)
	if err != nil {
		t.Error(err)
	}

	qg := NewESQueryGenerator("http://localhost")

	esquery := qg.Parse("products", "testing", req)

	if esquery != "http://localhost/products/testing/_search?q=field1:1,field2:2&fields=id" {
		t.Errorf("The URL didn't match: %s", esquery)
	}
}
