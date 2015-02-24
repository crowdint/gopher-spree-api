package repositories

import (
	"reflect"
	"testing"
)

func TestOptionTypeRepo(t *testing.T) {
	err := InitDB(true)

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	optionTypeRepo := NewOptionTypeRepo()

	optionTypes, err := optionTypeRepo.FindByProductIds([]int64{3})

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	not := len(optionTypes)

	if not < 1 {
		t.Error("Invalid number of rows: %d", not)
		return
	}

	temp := reflect.ValueOf(*optionTypes[0]).Type().String()

	if temp != "models.OptionType" {
		t.Error("Invalid type", t)
	}
}
