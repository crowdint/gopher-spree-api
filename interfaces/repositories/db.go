package repositories

import (
	"errors"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Spree_db *gorm.DB

type DbRepo struct {
	dbHandler *gorm.DB
}

const (
	DbUrlEnvName       = "DATABASE_URL"
	DbEngineEnvName    = "DATABASE_ENGINE"
	MaxIdleConnections = "MAX_IDLE_CONNS"
	MaxOpenConnections = "MAX_OPEN_CONNS"
)

func InitDB() error {
	dbUrl := os.Getenv(DbUrlEnvName)
	dbEngine := os.Getenv(DbEngineEnvName)

	if dbUrl == "" {
		return errors.New(DbUrlEnvName + " environment variable not found")
	}

	if dbEngine == "" {
		return errors.New(DbEngineEnvName + " environment variable not found")
	}

	db, err := gorm.Open(dbEngine, dbUrl)

	dbLog, _ := strconv.ParseBool(os.Getenv("DATABASE_DEBUG"))
	db.LogMode(dbLog)

	if err != nil {
		return err
	}

	maxIdle := os.Getenv(MaxIdleConnections)
	db.DB().SetMaxIdleConns(getIntegerOrDefault(maxIdle, 10))

	maxOpen := os.Getenv(MaxOpenConnections)
	db.DB().SetMaxOpenConns(getIntegerOrDefault(maxOpen, 100))

	db.SingularTable(true)

	Spree_db = &db

	return nil
}

func getIntegerOrDefault(value string, def int) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return number
}

func NewDatabaseRepository() *DbRepo {
	return &DbRepo{Spree_db}
}

func (this *DbRepo) FindBy(model interface{}, attrs map[string]interface{}) error {
	return this.dbHandler.First(model, attrs).Error
}
