package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"

	"testing"
)

func TestOptionTypeInteractor_GetJsonOptionTypesMap(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	optionTypeInteractor := NewOptionTypeInteractor()

	optionTypeMap, err := optionTypeInteractor.GetJsonOptionTypesMap([]int64{3})

	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	noptionTypes := len(optionTypeMap)

	if noptionTypes < 1 {
		t.Errorf("Wrong number of records %d", noptionTypes)
	}

}

func TestOptionTypeInteractor_modelsToJsonOptionTypesMap(t *testing.T) {
	optionTypeslice := []*domain.OptionType{
		&domain.OptionType{
			Id:           1,
			Name:         "tshirt-size",
			Presentation: "Size",
			Position:     1,
			ProductId:    3,
		},
		&domain.OptionType{
			Id:           2,
			Name:         "tshirt-color",
			Presentation: "Color",
			Position:     2,
			ProductId:    3,
		},
	}

	optionTypeInteractor := NewOptionTypeInteractor()

	jsonOptionTypeMap := optionTypeInteractor.modelsToJsonOptionTypesMap(optionTypeslice)

	optionType1 := jsonOptionTypeMap[3][0]
	optionType2 := jsonOptionTypeMap[3][1]

	if optionType1 == nil || optionType2 == nil {
		t.Error("Error: nil value on map")
	}

	if optionType1.Id != 1 || optionType1.Name != "tshirt-size" || optionType1.Presentation != "Size" || optionType1.Position != 1 {
		t.Error("Invalid values for first option type")
	}

	if optionType2.Id != 2 || optionType2.Name != "tshirt-color" || optionType2.Presentation != "Color" || optionType2.Position != 2 {
		t.Error("Invalid values for second option type")
	}

}

func TestOptionTypeInteractor_toJson(t *testing.T) {
	optionType := &domain.OptionType{
		Id:           2,
		Name:         "tshirt-color",
		Presentation: "Color",
		Position:     2,
		ProductId:    3,
	}

	optionTypeInteractor := NewOptionTypeInteractor()

	jsonOptionType := optionTypeInteractor.toJson(optionType)

	if jsonOptionType.Id != 2 || jsonOptionType.Name != "tshirt-color" || jsonOptionType.Presentation != "Color" || jsonOptionType.Position != 2 {
		t.Error("Invalid values for first option type")
	}
}
