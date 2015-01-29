package json

import (
	"encoding/json"
	"testing"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductInteractor_GetMergedResponse(t *testing.T) {
	repositories.InitDB()

	defer repositories.Spree_db.Close()

	productInteractor := NewProductInteractor()

	jsonProductSlice, err := productInteractor.GetMergedResponse()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if len(jsonProductSlice) < 1 {
		t.Error("Error: Invalid number of rows")
	}

	jsonBytes, err := json.Marshal(jsonProductSlice)
	if err != nil {
		t.Error("Error:", err.Error())
	}

	t.Error(string(jsonBytes))
}
