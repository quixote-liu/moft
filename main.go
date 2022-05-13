package main

import (
	"log"
	"moft/database"
	"moft/handler"
	"net"
	"net/http"

	"github.com/quixote-liu/config"
)

var conf = config.CONF()

func main() {
	// parse configuration.
	if err := conf.LoadConfiguration("./config.conf"); err != nil {
		log.Printf("parse configuration failed: %v", err)
		return
	}

	// init database.
	db, err := database.InitDatabase()
	if err != nil {
		log.Printf("init database failed: %v", err)
		return
	}
	if err := database.MigrateTables(db); err != nil {
		log.Printf("migrate databases tables failed: %v", err)
		return
	}

	// register routers.
	http.HandleFunc("/api/v1/file/receive", handler.ReceiveFile)

	host := conf.GetString("system", "host")
	port := conf.GetString("system", "port")
	server := net.JoinHostPort(host, port)
	log.Printf("start server on %s", server)
	log.Println(http.ListenAndServe(server, nil))
}
