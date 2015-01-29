package repositories

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/crowdint/gopher-spree-api/configs"
)

var Spree_db *gorm.DB

type DbRepo struct {
	dbHandler *gorm.DB
}

func InitDB() error {
	dbUrl := configs.Get(configs.DB_URL)
	dbEngine := configs.Get(configs.DB_ENGINE)

	if dbUrl == "" {
		return errors.New(configs.DB_URL + " environment variable not found")
	}

	if dbEngine == "" {
		return errors.New(configs.DB_ENGINE + " environment variable not found")
	}

	db, err := gorm.Open(dbEngine, dbUrl)

	dbLog, _ := strconv.ParseBool(configs.Get(configs.DB_DEBUG))
	db.LogMode(dbLog)

	if err != nil {
		return err
	}

	maxIdle := configs.Get(configs.DB_IDLE_CONNS)
	db.DB().SetMaxIdleConns(getIntegerOrDefault(maxIdle, 10))

	maxOpen := configs.Get(configs.DB_OPEN_CONNS)
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
