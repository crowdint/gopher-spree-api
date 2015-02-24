package json

import (
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"testing"
)

func TestClassificationInteractor_GetJsonClassificationsMap(t *testing.T) {
	err := repositories.InitDB(true)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	classificationInteractor := NewClassificationInteractor()

	classificationMap, err := classificationInteractor.GetJsonClassificationsMap([]int64{1, 2, 3})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	nclassifications := len(classificationMap)

	if nclassifications < 1 {
		t.Errorf("Wrong number of records %d", nclassifications)
	}

}
