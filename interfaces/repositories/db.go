package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var spree_db *gorm.DB

type DbRepo struct {
	dbHandler *gorm.DB
}

func InitDB() error {
	db, err := gorm.Open("postgres", "dbname=spree_dev sslmode=disable")
	if err != nil {
		return err
	}

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	db.SingularTable(true)

	spree_db = &db

	return nil
}
