package main

import (
	"cryptoViewer/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/random", handlers.RandomHandler).Methods("GET")
	r.HandleFunc("/random", handlers.RandomHandlerError).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
