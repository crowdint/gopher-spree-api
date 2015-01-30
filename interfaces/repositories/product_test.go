package repositories

import (
	"reflect"
	"testing"
)

func TestProductRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	productRepo := NewProductRepo()

	product, err := productRepo.FindById(1)

	if product.Name == "" {
		t.Error("No name found")
	}

	if product.CreatedAt.IsZero() {
		t.Error("No created_at found")
	}

	if err != nil {
		t.Error("dbHandler error:", err)
	}

	productSlice, err := productRepo.List(1, 10)
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	nv := len(productSlice)

	if nv > 10 || nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
		return
	}

	temp := reflect.ValueOf(*productSlice[0]).Type().String()

	if temp != "models.Product" {
		t.Error("Invalid type", t)
	}

	if err != nil {
		t.Error("dbHandler error:", err)
	}

	count, err := productRepo.CountAll()
	if err != nil {
		t.Error("dbHandler error:", err)
	}

	if count < 1 {
		t.Error("dbHandler error:", err)
	}
}
