package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/json"
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

	if temp != "json.OptionValue" {
		t.Error("Invalid type", t)
	}
}

func TestOptionValueRepository_AllByVariantAssociation(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	optionValueRepo := NewOptionValueRepo()
	variant := &json.Variant{Id: 17}
	optionValues := optionValueRepo.AllByVariantAssociation(variant)

	if len(optionValues) < 1 {
		t.Errorf("There aren't option values from this variant: %d", variant.Id)
	}
}
