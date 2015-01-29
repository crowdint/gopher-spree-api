package repositories

import (
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	os.Setenv(DbUrlEnvName, "dbname=spree_dev sslmode=disable")
	os.Setenv(DbEngineEnvName, "postgres")

	err := InitDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}

	Spree_db.Close()
}
