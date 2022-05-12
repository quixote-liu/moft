package main

import (
	"log"
	"moft/handler"
	"moft/model"
	"net"
	"net/http"
)

func main() {
	// parse configuration.
	if err := model.CONF.LoadConfiguration("./config.conf"); err != nil {
		log.Printf("parse configuration failed: %v", err)
		return
	}
	conf := model.CONF

	// register routers.
	http.HandleFunc("/api/v1/file/receive", handler.ReceiveFile)

	host := conf.GetString("system", "host")
	port := conf.GetString("system", "port")
	server := net.JoinHostPort(host, port)
	log.Printf("start server on %s", server)
	log.Println(http.ListenAndServe(server, nil))
}
