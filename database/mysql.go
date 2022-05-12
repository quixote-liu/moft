package database

import (
	"fmt"
)

type mysqlOptions struct {
	username string
	password string
	port     string
	host     string
	dbname   string
}

func (opts mysqlOptions) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		opts.username, opts.password, opts.host, opts.port, opts.dbname)
}

func mysqlOpts() mysqlOptions {
	return mysqlOptions{
		username: conf.GetString("database", "username"),
		password: conf.GetString("database", "password"),
		port:     conf.GetString("database", "port"),
		host:     conf.GetString("database", "host"),
		dbname:   conf.GetString("database", "dbname"),
	}
}
