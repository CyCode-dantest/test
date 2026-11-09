package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "project": "project3"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods(http.MethodGet)
	http.ListenAndServe(":9090", r)
}
