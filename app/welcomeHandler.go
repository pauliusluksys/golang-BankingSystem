package app

import (
	"encoding/json"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	data := "welcome to my humble website"
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
