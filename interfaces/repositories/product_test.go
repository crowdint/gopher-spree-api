package repositories

import (
	"os"
	"reflect"
	"testing"
)

func TestProductRepo(t *testing.T) {
	os.Setenv(dbUrlEnvName, "dbname=spree_dev sslmode=disable")
	os.Setenv(dbmsEnvName, "postgres")

	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer spree_db.Close()

	productRepo := NewProductRepo()

	product := productRepo.FindById(1)

	if product.Name == "" {
		t.Error("No name found")
	}

	if product.CreatedAt.IsZero() {
		t.Error("No created_at found")
	}

	productSlice, err := productRepo.List()
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	nv := len(productSlice)

	if nv < 1 {
		t.Errorf("Invalid number of rows: %d", nv)
	}

	temp := reflect.ValueOf(productSlice[0]).Type().String()

	if temp != "models.Product" {
		t.Error("Invalid type", t)
	}
}
