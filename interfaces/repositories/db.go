package repositories

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/crowdint/gopher-spree-api/configs"
	gsk "github.com/crowdint/gransak"
)

var Spree_db *gorm.DB

type DbRepository struct {
	dbHandler *gorm.DB
}

type Not struct {
	Key    string
	Values []interface{}
}

func InitDB() error {
	dbUrl := configs.Get(configs.DB_URL)
	dbEngine := configs.Get(configs.DB_ENGINE)

	if dbEngine == "postgres" {
		//By default it uses MySQL
		gsk.Gransak.SetEngine("postgresql")
	}

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

func NewDatabaseRepository() *DbRepository {
	return &DbRepository{Spree_db}
}

func extractPaginationValues(attrs map[string]interface{}) (limit, offset int) {
	if attrs["limit"] != nil && attrs["offset"] != nil {
		limit = attrs["limit"].(int)
		delete(attrs, "limit")
		offset = (attrs["offset"].(int) - 1) * limit
		delete(attrs, "offset")
	}

	return
}

func getIntegerOrDefault(value string, def int) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return number
}

func (this *DbRepository) All(collection interface{}, options map[string]interface{}, query interface{}, values ...interface{}) error {
	limit, offset := extractPaginationValues(options)
	dbHandler := this.dbHandler

	if limit > 0 {
		dbHandler = dbHandler.Limit(limit).Offset(offset)
	}

	dbHandler = orderByIfPresent(dbHandler, options)
	dbHandler = notIfPresent(dbHandler, options)

	return dbHandler.Where(query, values...).Find(collection).Error
}

func (this *DbRepository) Association(model interface{}, association interface{}, attribute string) {
	this.dbHandler.Model(model).Related(association, attribute)
}

func (this *DbRepository) Count(model interface{}, query string, params []interface{}) (count int64, err error) {
	err = this.dbHandler.Model(model).Where(query, params).Count(&count).Error
	return
}

func (this *DbRepository) FindBy(model interface{}, options map[string]interface{}, where map[string]interface{}) error {
	dbHandler := this.dbHandler
	dbHandler = notIfPresent(dbHandler, options)
	return dbHandler.First(model, where).Error
}

func orderByIfPresent(dbHandler *gorm.DB, options map[string]interface{}) *gorm.DB {
	if options["order"] != nil {
		orderBy := options["order"].(string)
		if orderBy != "" {
			dbHandler = dbHandler.Order(orderBy)
		}
	}
	return dbHandler
}

func notIfPresent(dbHandler *gorm.DB, options map[string]interface{}) *gorm.DB {
	if options["not"] != nil {
		not := options["not"].(Not)
		dbHandler = dbHandler.Not(not.Key, not.Values)
	}
	return dbHandler
}
