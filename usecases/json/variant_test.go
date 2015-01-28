package json

import (
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestVariantInteractor_GetJsonVariantsMap(t *testing.T) {
	repositories.InitDB()

	defer repositories.spree_db.Close()

	variantInteractor := NewVariantInteractor()

	variantMap := variantInteractor.GetJsonVariantsMap([]int64{1, 2, 3})

	nvariants := len(variantMap)

	if nvariants != 3 {
		t.Errorf("Wrong number of records %d", nvariants)
	}
}

func TestVariantInteractor_modelsToJsonVariantsMap(t *testing.T) {
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

	jsonVariantMap := variantInteractor.modelsToJsonVariantsMap(variantSlice)

	v1 := jsonVariantMap[10]
	v2 := jsonVariantMap[20]

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
