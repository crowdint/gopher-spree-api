package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"testing"
)

func TestOptionValueInteractor_GetJsonOptionValuesMap(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	optionValue := &domain.OptionValue{
		Id:           10,
		OptionTypeId: 1,
	}

	optionType := &domain.OptionType{
		Id: 1,
	}

	repositories.Spree_db.Create(optionValue)
	repositories.Spree_db.Create(optionType)
	repositories.Spree_db.Exec("INSERT INTO spree_option_values_variants(option_value_id, variant_id) values(10, 17)")

	optionValueInteractor := NewOptionValueInteractor()

	optionValueMap, err :=
		optionValueInteractor.GetJsonOptionValuesMap([]int64{17})

	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	noptionValues := len(optionValueMap)

	if noptionValues < 1 {
		t.Errorf("Wrong number of records %d", noptionValues)
	}

}

func TestOptionValueInteractor_modelsToJsonOptionValuesMap(t *testing.T) {
	optionValueslice := []*domain.OptionValue{
		&domain.OptionValue{
			Id:                     2,
			Name:                   "Medium",
			Presentation:           "M",
			VariantId:              17,
			OptionTypeName:         "thshirt-size",
			OptionTypeId:           1,
			OptionTypePresentation: "Size",
		},
	}

	optionValueInteractor := NewOptionValueInteractor()

	jsonOptionValueMap := optionValueInteractor.modelsToJsonOptionValuesMap(optionValueslice)

	optionValue := jsonOptionValueMap[17][0]

	if optionValue.Id != 2 || optionValue.Name != "Medium" || optionValue.Presentation != "M" {
		t.Error("Invalid values for first option type")
	}

}
