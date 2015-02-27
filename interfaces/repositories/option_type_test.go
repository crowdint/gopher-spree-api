package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestOptionTypeRepo(t *testing.T) {
	err := InitDB(true)

	defer func() {
		Spree_db.Rollback()
		Spree_db.Close()
	}()

	optionType := &domain.OptionType{Id: 1}

	Spree_db.Create(optionType)
	Spree_db.Exec("INSERT into spree_product_option_types(product_id, option_type_id) values(3, 1)")

	if err != nil {
		t.Error("An error has ocurred", err)
		return
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
		return
	}

	optionTypeRepo := NewOptionTypeRepo()

	optionTypes, err := optionTypeRepo.FindByProductIds([]int64{3})

	if err != nil {
		t.Error("An error has ocurred", err)
		return
	}

	not := len(optionTypes)

	if not < 1 {
		t.Error("Invalid number of rows: %d", not)
		return
	}

	temp := reflect.ValueOf(*optionTypes[0]).Type().String()

	if temp != "domain.OptionType" {
		t.Error("Invalid type", t)
	}
}
