package json

import (
	"encoding/json"
	"testing"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestResponseInteractor(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	productInteractor := NewProductInteractor()

	interactor := SpreeResponseFetcher

	params := newDummyResponseParams(2, 0, "", nil)

	response, err := interactor.GetResponse(productInteractor, params)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	jsonBytes, err := json.Marshal(response)

	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
		return
	}

	strJson := string(jsonBytes)

	if strJson == "" {
		t.Error("Error: Empty json string")
		return
	}
}
