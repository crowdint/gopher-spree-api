package json

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"

	"testing"
)

func TestVariantInteractor_GetJsonVariantsMap(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	variantInteractor := NewVariantInteractor()

	variantMap, err := variantInteractor.GetJsonVariantsMap([]int64{1, 2, 3})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	nvariants := len(variantMap)

	if nvariants < 1 {
		t.Errorf("Wrong number of records %d", nvariants)
	}

	varray1 := variantMap[1]
	if len(varray1) < 1 {
		t.Error("No variants found")
	}

	v1 := varray1[0]

	if v1.Price == "" {
		t.Error("Variant has no price")
	}
}

func TestVariantInteractor_modelsToJsonVariantsMap(t *testing.T) {
	err := repositories.InitDB()
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	defer repositories.Spree_db.Close()

	variantSlice := []*models.Variant{
		&models.Variant{
			Id:        1,
			Sku:       "sku0001",
			Price:     "9.99",
			ProductId: 10,
		},
		&models.Variant{
			Id:        2,
			Sku:       "sku0002",
			Price:     "10.99",
			ProductId: 20,
		},
	}

	variantInteractor := NewVariantInteractor()

	jsonVariantMap, err := variantInteractor.modelsToJsonVariantsMap(variantSlice)

	if err != nil {
		t.Error("Error: something went wrong", err.Error)
	}

	v1 := jsonVariantMap[10][0]
	v2 := jsonVariantMap[20][0]

	if v1 == nil || v2 == nil {
		t.Error("Error: nil value on map")
	}

	if v1.Id != 1 || v1.Sku != "sku0001" || v1.Price != "9.99" {
		t.Error("Invalid values for first variant")
	}

	if v2.Id != 2 || v2.Sku != "sku0002" || v2.Price != "10.99" {
		t.Error("Invalid values for second variant")
	}
}

func TestVariantInteractor_toJson(t *testing.T) {
	variant := &models.Variant{
		Id:        1,
		Sku:       "sku0001",
		Price:     "9.99",
		ProductId: 10,
	}

	variantInteractor := NewVariantInteractor()

	jsonVariant := variantInteractor.toJson(variant)

	if jsonVariant.Id != 1 || jsonVariant.Sku != "sku0001" || jsonVariant.Price != "9.99" {
		t.Error("Invalid values for second json.Variant")
	}
}
