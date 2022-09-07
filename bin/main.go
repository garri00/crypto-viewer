package main

import (
	"crypto-viewer/scr/handlers"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	fmt.Println("server start")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", handlers.Coins)

	http.ListenAndServe(":8080", r)
	fmt.Println("server STOP")

}
