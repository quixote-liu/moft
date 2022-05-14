package main

import (
	"fmt"
	"moft/model"
	"moft/util"
	"net/http"
)

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
	if h == nil {
		util.JSONResponse(w, http.StatusNotFound, model.H{
			"error":  "router not found",
			"method": method,
			"path":   pattern,
		})
		return
	}
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
