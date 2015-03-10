package api

import (
	"net/http"
	"testing"
)

func TestESFetcher_GetProductIds(t *testing.T) {
	apiUrl := "http://localhost:1440"
	req, err := http.NewRequest("GET", apiUrl+"?name=Spree", nil)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	_, err = esfetcher.GetProducIds("test", "product", req)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}
}
