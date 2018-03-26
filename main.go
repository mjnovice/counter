package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Place struct {
	Name  string `json:"name,omitempty"`
	Count int64  `json:"count,omitempty"`
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	place := Place{Name: "Plaza Singapura", Count: 123}
	json.NewEncoder(w).Encode(place)
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/place/{name}", GetPlace).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
