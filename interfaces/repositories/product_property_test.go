package repositories

import (
	"reflect"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestProductPropertyRepo(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	Spree_db.Create(&domain.ProductProperty{Id: 1, ProductId: 1, PropertyId: 1})
	Spree_db.Exec("INSERT INTO spree_properties(id, presentation) values(1, 'foo')")

	productPropertyRepo := NewProductPropertyRepo()

	productProperties, err := productPropertyRepo.FindByProductIds([]int64{1, 2})

	if err != nil {
		t.Error("An error has ocurred:", err)
	}

	npp := len(productProperties)

	if npp < 1 {
		t.Errorf("Invalid number of rows: %d", npp)
		return
	}

	temp := reflect.ValueOf(*productProperties[0]).Type().String()

	if temp != "domain.ProductProperty" {
		t.Error("Invalid type", t)
	}

}
