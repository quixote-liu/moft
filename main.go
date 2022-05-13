package main

import (
	"fmt"
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

	// create server mux.
	mux := NewServerMux()

	// register routers.
	mux.HandleFunc("/api/v1/file/receivce", handler.ReceiveFile)

	// register users.
	userHandler := handler.NewUserHandler(db)
	{
		mux.POST("/api/v1/user/register", userHandler.Register)
	}

	host := conf.GetString("system", "host")
	port := conf.GetString("system", "port")
	server := net.JoinHostPort(host, port)
	log.Printf("start server on %s", server)
	log.Println(http.ListenAndServe(server, mux))
}

type ServerMux struct {
	*http.ServeMux
	routers map[string]map[string]http.HandlerFunc
}

func NewServerMux() *ServerMux {
	return &ServerMux{
		ServeMux: http.NewServeMux(),
		routers:  make(map[string]map[string]http.HandlerFunc),
	}
}

func (mux *ServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Path
	method := r.Method
	h := mux.findHandlerFunc(method, pattern)
	h(w, r)
}

func (mux *ServerMux) POST(pattern string, h http.HandlerFunc) {
	mux.registerHandlerFunc(http.MethodPost, pattern, h)
}

func (mux *ServerMux) GET(pattern string, h http.HandlerFunc) {
	mux.registerHandlerFunc(http.MethodGet, pattern, h)
}

func (mux *ServerMux) PUT(pattern string, h http.HandlerFunc) {
	mux.registerHandlerFunc(http.MethodPut, pattern, h)
}

func (mux *ServerMux) DELETE(pattern string, h http.HandlerFunc) {
	mux.registerHandlerFunc(http.MethodDelete, pattern, h)
}

func (mux *ServerMux) findHandlerFunc(method, pattern string) http.HandlerFunc {
	pp, ok := mux.routers[method]
	if !ok {
		return nil
	}
	h, ok := pp[pattern]
	if !ok {
		return nil
	}
	return h
}

func (mux *ServerMux) registerHandlerFunc(method, pattern string, h http.HandlerFunc) {
	if _, ok := mux.routers[method]; !ok {
		mux.routers[method] = make(map[string]http.HandlerFunc)
	}
	if _, ok := mux.routers[method][pattern]; ok {
		err := fmt.Errorf("the router conflict: %v", pattern)
		panic(err)
	}

	mux.routers[method][pattern] = h
}
