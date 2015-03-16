package spree

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/utils"
)

const (
	COUNTRY_ID             = "spree/app_configuration/default_country_id"
	CURRENCY               = "spree/app_configuration/currency"
	API_AUTHENTICATION     = "spree/api_configuration/requires_authentication"
	TRACK_INVENTORY_LEVELS = "spree/app_configuration/track_inventory_levels"
	MASTER_PRICE           = "spree/app_configuration/require_master_price"
)

var (
	dbHandler   *gorm.DB
	spreeConfig = map[string]string{
		COUNTRY_ID:             getDbOrDefault(COUNTRY_ID, "232"),
		CURRENCY:               getDbOrDefault(CURRENCY, "USD"),
		API_AUTHENTICATION:     getDbOrDefault(API_AUTHENTICATION, "true"),
		TRACK_INVENTORY_LEVELS: getDbOrDefault(TRACK_INVENTORY_LEVELS, "true"),
		MASTER_PRICE:           getDbOrDefault(MASTER_PRICE, "true"),
	}
)

type Preference struct {
	Id    int64
	Value string
	Key   string
}

func (p Preference) TableName() string {
	return "spree_preferences"
}

func Get(key string) string {
	return spreeConfig[key]
}

func Set(key, value string) {
	spreeConfig[key] = value
}

func SetAndSave(key, value string) error {
	p, err := findPreferenceByKey(key)
	if err != nil {
		return err
	}
	p.Value = serialize(value)
	err = dbHandler.Save(p).Error
	if err != nil {
		return err
	}

	Set(key, value)
	return nil
}

func IsInventoryTrackingEnabled() bool {
	value, err := strconv.ParseBool(spreeConfig[TRACK_INVENTORY_LEVELS])

	if err != nil {
		return true
	}

	return value
}

func IsAuthenticationRequired() bool {
	value, err := strconv.ParseBool(spreeConfig[API_AUTHENTICATION])
	if err != nil {
		return true // because by default authentication is true in spree
	}

	return value
}

func getDbOrDefault(key string, defaultValue string) string {
	if err := initDbRepo(); err != nil {
		return defaultValue
	}

	p, err := findPreferenceByKey(key)
	if err != nil {
		return defaultValue
	}

	p.Value = unserialize(p.Value)
	return p.Value
}

func findPreferenceByKey(key string) (*Preference, error) {
	p := &Preference{}
	err := dbHandler.First(p, map[string]interface{}{"key": key}).Error
	return p, err
}

func initDbRepo() error {
	if dbHandler == nil {
		dbUrl := configs.Get(configs.DB_URL)
		dbEngine := configs.Get(configs.DB_ENGINE)
		utils.LogrusInfo("initDbRepo", "Initializing database")
		db, err := gorm.Open(dbEngine, dbUrl)
		if err != nil {
			return err
		}

		dbLog, _ := strconv.ParseBool(configs.Get(configs.DB_DEBUG))
		db.LogMode(dbLog)
		dbHandler = &db
	}

	return nil
}

func unserialize(s string) string {
	s = strings.Replace(s, "--- ", "", -1)
	return strings.Replace(s, "\n...\n", "", -1)
}

func serialize(s string) string {
	return "--- " + s + "\n...\n"

}
