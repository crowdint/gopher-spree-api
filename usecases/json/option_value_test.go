package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"testing"
)

func TestOptionValueInteractor_GetJsonOptionValuesMap(t *testing.T) {
	err := repositories.InitDB(true)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

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
