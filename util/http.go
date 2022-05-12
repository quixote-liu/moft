package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJSONErr(w http.ResponseWriter, code int, msg interface{}) {
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

func BindingJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
