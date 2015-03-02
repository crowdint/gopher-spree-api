package json

import (
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"

	"testing"
)

func TestVariantInteractor_GetJsonVariantsMap(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	repositories.Spree_db.Create(&domain.Variant{Id: 1, ProductId: 1, CostPrice: "10"})
	repositories.Spree_db.Exec("INSERT INTO spree_stock_items(variant_id) values(1)")
	repositories.Spree_db.Exec("INSERT INTO spree_prices(variant_id, currency) values(1, 'USD')")

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
}

func TestVariantInteractor_modelsToJsonVariantsMap(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	variantSlice := []*domain.Variant{
		&domain.Variant{
			Id:        1,
			Sku:       "sku0001",
			Price:     9.99,
			ProductId: 10,
		},
		&domain.Variant{
			Id:        2,
			Sku:       "sku0002",
			Price:     10.99,
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

	if v1.Id != 1 || v1.Sku != "sku0001" || v1.Price != 9.99 {
		t.Error("Invalid values for first variant")
	}

	if v2.Id != 2 || v2.Sku != "sku0002" || v2.Price != 10.99 {
		t.Error("Invalid values for second variant")
	}
}

func TestVariantInteractor_toJson(t *testing.T) {
	variant := &domain.Variant{
		Id:        1,
		Sku:       "sku0001",
		Price:     9.99,
		ProductId: 10,
	}

	variantInteractor := NewVariantInteractor()

	jsonVariant := variantInteractor.toJson(variant)

	if jsonVariant.Id != 1 || jsonVariant.Sku != "sku0001" || jsonVariant.Price != 9.99 {
		t.Error("Invalid values for second domain.Variant")
	}
}
