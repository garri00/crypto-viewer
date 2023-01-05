package api

import (
	"net/http"

	"crypto-viewer/api/handlers"
	"crypto-viewer/api/handlers/adapters"
	"crypto-viewer/api/handlers/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("welcome"))
		w.WriteHeader(http.StatusOK)
	})

	adapter := adapters.NewCoins(resty.New())
	usecase := usecases.NewCoins(adapter)
	c := handlers.NewRestyClient(usecase)

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", c.CoinsResty)
	return r
}
