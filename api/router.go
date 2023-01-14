package api

import (
	"crypto-viewer/api/adapters"
	"crypto-viewer/api/usecases"
	"net/http"

	"crypto-viewer/api/handlers"
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

	restyClient := resty.New()
	coinsAdapter := adapters.NewCoins(restyClient)
	exchangeAdapter := adapters.NewExchange(restyClient)

	coinsUseCase := usecases.NewCoins(coinsAdapter, exchangeAdapter)
	saveDataUseCase := usecases.NewSaveData()

	c := handlers.CoinsHendler(coinsUseCase, saveDataUseCase)

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", c.CoinsResty)
	return r
}
