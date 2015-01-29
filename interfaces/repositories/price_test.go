package repositories

import (
	"os"
	"reflect"
	"testing"
)

func TestPriceRepo(t *testing.T) {
	os.Setenv(DbUrlEnvName, "dbname=spree_dev sslmode=disable")
	os.Setenv(DbEngineEnvName, "postgres")

	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	defer Spree_db.Close()

	priceRepo := NewPriceRepo()

	price := priceRepo.GetByVariant(26)

	temp := reflect.ValueOf(price).Type().String()

	if temp != "models.Price" {
		t.Error("Invalid type", t)
	}

	currency := price.Currency

	if currency != "USD" {
		t.Error("Wrong currency price", t)
	}

}
