package database

import (
	"fmt"
	"moft/model"

	"github.com/quixote-liu/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conf = config.CONF()

func InitDatabase() (*gorm.DB, error) {
	dbType := conf.GetString("system", "database_type")

	var db *gorm.DB
	var err error

	switch dbType {
	case "mysql":
		dsn := mysqlOpts().dsn()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		err = fmt.Errorf("the database type %s does not support", dbType)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Ticket{},
	)
}
