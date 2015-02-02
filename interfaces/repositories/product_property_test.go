package repositories

import (
	"reflect"
	"testing"
)

func TestProductPropertyRepo(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	productPropertyRepo := NewProductPropertyRepo()

	productProperties, err := productPropertyRepo.FindByProductIds([]int64{1, 2})

	if err != nil {
		t.Error("waka")
	}

	npp := len(productProperties)

	if npp < 1 {
		t.Errorf("Invalid number of rows: %d", npp)
		return
	}

	temp := reflect.ValueOf(*productProperties[0]).Type().String()

	if temp != "models.ProductProperty" {
		t.Error("Invalid type", t)
	}

}
