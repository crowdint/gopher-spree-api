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

type DbRepo struct {
	dbHandler *gorm.DB
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

func NewDatabaseRepository() *DbRepo {
	return &DbRepo{Spree_db}
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

func (this *DbRepo) All(collection interface{}, options map[string]interface{}, query interface{}, values ...interface{}) error {
	limit, offset := extractPaginationValues(options)

	if limit == 0 {
		return this.dbHandler.Where(query, values...).Find(collection).Error
	}

	return this.dbHandler.Offset(offset).Limit(limit).Where(query, values...).Find(collection).Error
}

func (this *DbRepo) Association(model interface{}, association interface{}, attribute string) {
	this.dbHandler.Model(model).Related(association, attribute)
}

func (this *DbRepo) Count(model interface{}, query string, params []interface{}) (count int64, err error) {
	err = this.dbHandler.Model(model).Where(query, params).Count(&count).Error
	return
}

func (this *DbRepo) FindBy(model interface{}, attrs map[string]interface{}) error {
	return this.dbHandler.First(model, attrs).Error
}

func (this *DbRepo) SumLineItemsQuantityByOrderIds(ids []int64) (map[int64]int64, error) {
	orderQuantities := []struct {
		Id  int64
		Sum int64
	}{}

	err := this.dbHandler.
		Table("spree_line_items").
		Select("order_id AS id, SUM(quantity) AS sum").
		Where("order_id IN (?)", ids).
		Group("order_id").
		Scan(&orderQuantities).
		Error

	orderQuantitiesMap := map[int64]int64{}
	for _, oq := range orderQuantities {
		orderQuantitiesMap[oq.Id] = oq.Sum
	}

	return orderQuantitiesMap, err
}
