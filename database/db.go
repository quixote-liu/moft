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

	var dsn string

	switch dbType {
	case "mysql":
		dsn = mysqlOpts().dsn()
	default:
		return nil, fmt.Errorf("the database type %s does not support", dbType)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := migrateTables(db); err != nil {
		return nil, fmt.Errorf("migrate tables failed: %v", err)
	}

	return db, nil
}

func migrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Ticket{},
	)
}
