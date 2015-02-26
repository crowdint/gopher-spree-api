package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
)

func TestOptionValueRepo(t *testing.T) {
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	optionValue := &models.OptionValue{
		Id:           1,
		OptionTypeId: 1,
	}

	optionType := &models.OptionType{
		Id: 1,
	}

	Spree_db.Create(optionValue)
	Spree_db.Create(optionType)
	Spree_db.Exec("INSERT INTO spree_option_values_variants(option_value_id, variant_id) values(1, 17)")

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

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

	if temp != "domain.OptionValue" {
		t.Error("Invalid type", t)
	}
}

func TestOptionValueRepository_AllByVariantAssociation(t *testing.T) {
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	optionValue := &models.OptionValue{
		Id: 1,
	}

	Spree_db.Create(optionValue)
	Spree_db.Exec("INSERT INTO spree_option_values_variants(option_value_id, variant_id) values(1, 17)")

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	optionValueRepo := NewOptionValueRepo()
	variant := &domain.Variant{Id: 17}
	optionValues := optionValueRepo.AllByVariantAssociation(variant)

	if len(optionValues) < 1 {
		t.Errorf("There aren't option values from this variant: %d", variant.Id)
	}
}
