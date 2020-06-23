package main

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

// PingHandler - Response for /ping request
func PingHandler(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("ping", "pong")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("auth", "true/false")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	r.HandleFunc("/auth", AuthHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
