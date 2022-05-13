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
	group := "mysql"
	return mysqlOptions{
		username: conf.GetString(group, "username"),
		password: conf.GetString(group, "password"),
		port:     conf.GetString(group, "port"),
		host:     conf.GetString(group, "host"),
		dbname:   conf.GetString(group, "dbname"),
	}
}
