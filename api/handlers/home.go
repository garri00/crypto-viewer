package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
}
