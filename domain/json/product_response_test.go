package json

import (
	"encoding/json"
	"testing"
)

func TestProductResponseStructure(t *testing.T) {
	expected := `{"products":[],"count":1,"pages":1,"current_page":1}`

	productResponse := ProductResponse{
		Products:    []*Product{},
		Count:       1,
		Pages:       1,
		CurrentPage: 1,
	}

	jsonBytes, err := json.Marshal(productResponse)
	if err != nil {
		t.Error("An error has ocurred:", err.Error())
	}

	current := string(jsonBytes)

	if current != expected {
		t.Error("Error: Json mismatch %s : %s", current, expected)
	}
}
