package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"

	"testing"
)

func TestProductPropertyInteractor_GetJsonProductPropertiesMap(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	productPropertyInteractor := NewProductPropertyInteractor()

	productPropertyMap, err := productPropertyInteractor.GetJsonProductPropertiesMap([]int64{1, 2})

	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	nProductProperties := len(productPropertyMap)

	if nProductProperties < 1 {
		t.Errorf("Wrong number of records %d", nProductProperties)
	}

	ppArray1 := productPropertyMap[1]

	if len(ppArray1) < 1 {
		t.Error("No productProperties found")
	}
}

func TestProductPropertyInteractor_modelsToJsonProductPropertiesMap(t *testing.T) {
	productPropertySlice := []*domain.ProductProperty{
		&domain.ProductProperty{
			Id:           66,
			ProductId:    10,
			PropertyId:   3,
			Value:        "Men's",
			PropertyName: "Gender",
		},
		&domain.ProductProperty{
			Id:           1,
			ProductId:    3,
			PropertyId:   1,
			Value:        "Wilson",
			PropertyName: "Manufacturer",
		},
	}

	productPropertyInteractor := NewProductPropertyInteractor()

	jsonProductPropertyMap := productPropertyInteractor.modelsToJsonProductPropertiesMap(productPropertySlice)

	pp1 := jsonProductPropertyMap[10][0]
	pp2 := jsonProductPropertyMap[3][0]

	if pp1 == nil || pp2 == nil {
		t.Error("Error: nil value on map")
	}

	if pp1.Id != 66 {
		t.Error("Invalid values for first productProperty")
	}

	if pp2.Id != 1 {
		t.Error("Invalid values for second productProperty")
	}
}
