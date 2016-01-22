package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/geo", geoHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatal(err)
	}
}

func geoHandler(w http.ResponseWriter, r *http.Request) {
	loc := new(Location)
	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		fmt.Fprintf(w, "Can't decode request body. Err: %v", err)
		return
	}

	log.Println(loc.str())
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

func (l *Location) str() string {
	return fmt.Sprintf("Lat: %f, Long: %f", l.Latitude, l.Longitude)
}
