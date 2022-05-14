package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, code int, msg interface{}) {
	if msg == nil {
		w.WriteHeader(code)
		return
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Printf("[ERROR]: marshal response data failed: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}

func Status(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Write(nil)
}

func JSONBinding(r *http.Request, v interface{}) error {
	if v == nil {
		return fmt.Errorf("the binding value is nil")
	}
	return json.NewDecoder(r.Body).Decode(v)
}
