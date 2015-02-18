package repositories

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/models"
)

func TestDB(t *testing.T) {
	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	Spree_db.Close()
}

func TestAll(t *testing.T) {
	InitDB()

	r := NewDatabaseRepository()

	p := &[]models.Product{}

	err := r.All(p, Q{"limit": 10, "offset": 1})

	if err != nil {
		t.Errorf("DB.All %s", err)
	}

	if len(*p) == 0 {
		t.Errorf("DB.All should have found results")
	}
}

func TestAllWithConditions(t *testing.T) {
	InitDB()

	r := NewDatabaseRepository()

	p := &[]models.Product{}

	err := r.All(p, Q{"id": 1, "limit": 10, "offset": 1})

	if err != nil {
		t.Errorf("DB.All %s", err)
	}

	if len(*p) != 1 {
		t.Errorf("DB.All should not have found more than one result")
	}
}

func TestFindBy(t *testing.T) {
	InitDB()

	r := NewDatabaseRepository()

	p := &models.Product{}

	err := r.FindBy(p, nil)

	if err != nil {
		t.Errorf("DB.All %s", err)
	}
}

func TestFindByWithConditions(t *testing.T) {
	InitDB()

	r := NewDatabaseRepository()

	p := &models.Product{}

	err := r.FindBy(p, Q{"id": 1})

	if err != nil {
		t.Errorf("DB.All %s", err)
	}
}
