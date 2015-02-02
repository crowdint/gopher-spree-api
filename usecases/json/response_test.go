package json

import (
	"encoding/json"
	"testing"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestResponseInteractor(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()
	interactor := NewResponseInteractor()

	response, err := interactor.GetResponse(2, 0)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	jsonBytes, err := json.Marshal(response)

	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	strJson := string(jsonBytes)

	if strJson == "" {
		t.Error("Error: Empty json string")
	}
}
