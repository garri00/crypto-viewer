package api

import (
	"crypto-viewer/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("welcome"))
		w.WriteHeader(http.StatusOK)
	})

	RestyClient := resty.New()
	c := handlers.NewRestyClient(RestyClient)

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", c.CoinsResty)
	return r
}
