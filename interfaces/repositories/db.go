package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"errors"
	"os"
	"strconv"
)

var spree_db *gorm.DB

type DbRepo struct {
	dbHandler *gorm.DB
}

const (
	dbUrlEnvName       = "DATABASE_URL"
	maxIdleConnections = "MAX_IDLE_CONNS"
	maxOpenConnections = "MAX_OPEN_CONNS"
)

func InitDB() error {
	dbUrl := os.Getenv(dbUrlEnvName)

	if dbUrl == "" {
		return errors.New(dbUrlEnvName + " environment variable not found")
	}

	db, err := gorm.Open("postgres", "dbname=spree_dev sslmode=disable")
	if err != nil {
		return err
	}

	maxIdle := os.Getenv(maxIdleConnections)
	db.DB().SetMaxIdleConns(getIntegerOrDefault(maxIdle, 10))

	maxOpen := os.Getenv(maxOpenConnections)
	db.DB().SetMaxOpenConns(getIntegerOrDefault(maxOpen, 100))

	db.SingularTable(true)

	spree_db = &db

	return nil
}

func getIntegerOrDefault(value string, def int) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return number
}
