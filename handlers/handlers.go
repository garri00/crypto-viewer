package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
}
func randomFunc(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	res := (rand.Intn(max-min+1) + min)
	return res
}
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ramdom number: %v", randomFunc(10, 20))
}
func RandomHandlerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
