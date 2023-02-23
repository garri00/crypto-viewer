package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"

	"crypto-viewer/api/adapters"
	"crypto-viewer/api/handlers"
	"crypto-viewer/api/usecases"
	"crypto-viewer/src/config"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	configs := config.GetConfig()

	restyClient := resty.New()
	coinsAdapter := adapters.NewCoins(restyClient, configs)
	exchangeAdapter := adapters.NewExchange(restyClient, configs)

	coinsUseCase := usecases.NewCoins(coinsAdapter, exchangeAdapter)
	saveDataUseCase := usecases.NewSaveData()

	c := handlers.CoinsHendler(coinsUseCase, saveDataUseCase)

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", c.CoinsResty)

	return r
}
