package main

import (
	"cryptoViewer/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"math/rand"
	"net/http"
	"time"
)

func randomFunc(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	res := (rand.Intn(max-min+1) + min)
	return res
}
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/random", handlers.RandomHandler)
	r.Post("/random", handlers.RandomHandlerError)

	http.ListenAndServe("localhost:8080", r)
}
