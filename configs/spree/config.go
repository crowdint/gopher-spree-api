package spree

import (
	"strconv"
	"strings"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

const (
	SPREE_COUNTRY_ID         = "spree/app_configuration/default_country_id"
	SPREE_CURRENCY           = "spree/app_configuration/currency"
	SPREE_API_AUTHENTICATION = "spree/api_configuration/requires_authentication"
	TRACK_INVENTORY_LEVELS   = "spree/app_configuration/track_inventory_levels"
)

var (
	dbRepo      *repositories.DbRepo
	spreeConfig = map[string]string{
		SPREE_COUNTRY_ID:         getDbOrDefault(SPREE_COUNTRY_ID, "232"),
		SPREE_CURRENCY:           getDbOrDefault(SPREE_CURRENCY, "USD"),
		SPREE_API_AUTHENTICATION: getDbOrDefault(SPREE_API_AUTHENTICATION, "true"),
		TRACK_INVENTORY_LEVELS:   getDbOrDefault(TRACK_INVENTORY_LEVELS, "true"),
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
	err = repositories.Spree_db.Save(p).Error
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
	value, err := strconv.ParseBool(spreeConfig[SPREE_API_AUTHENTICATION])
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
	err := dbRepo.FindBy(p, repositories.Q{"key": key})
	return p, err
}

func initDbRepo() error {
	if repositories.Spree_db == nil {
		if err := repositories.InitDB(); err != nil {
			return err
		}
	}

	if dbRepo == nil {
		dbRepo = repositories.NewDatabaseRepository()
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
