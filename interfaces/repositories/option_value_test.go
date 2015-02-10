package repositories

import (
	"reflect"
	"testing"
)

func TestOptionValueRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	optionValueRepo := NewOptionValueRepo()

	optionValues, err := optionValueRepo.FindByVariantIds([]int64{17})
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	nv := len(optionValues)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*optionValues[0]).Type().String()

	if temp != "models.OptionValue" {
		t.Error("Invalid type", t)
	}

}
