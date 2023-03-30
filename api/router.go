package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"

	"crypto-viewer/api/adapters"
	"crypto-viewer/api/handlers"
	"crypto-viewer/api/usecases"
	"crypto-viewer/pkg/clients/mongo"
	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
)

func NewRouter(configs config.Configs) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	mongoClient, err := mongo.NewClient(context.Background(), configs.MongoConf)
	if err != nil {

	}
	storage := mongo.NewStorage("coins", mongoClient)

	com1 := entities.Coin{
		ID:                0,
		Name:              "sa",
		Symbol:            "sa",
		Slug:              "as",
		NumMarketPairs:    0,
		DateAdded:         time.Time{},
		MaxSupply:         nil,
		CirculatingSupply: 0,
		TotalSupply:       0,
		Quote: entities.Quote{
			USD: entities.USD{
				Price:       1000,
				LastUpdated: time.Time{},
			},
		},
	}

	usr1, err := storage.Create(context.Background(), com1)
	if err != nil {
		fmt.Println(usr1, "111")
	}

	fmt.Println(usr1, "adwwad")

	restyClient := resty.New()
	coinsAdapter := adapters.NewCoins(restyClient, configs, logger.Log)
	exchangeAdapter := adapters.NewExchange(restyClient, configs, logger.Log)

	coinsUseCase := usecases.NewCoins(coinsAdapter, exchangeAdapter, logger.Log)
	saveDataUseCase := usecases.NewSaveData(logger.Log)

	c := handlers.CoinsHendler(coinsUseCase, saveDataUseCase, logger.Log)

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", c.CoinsResty)

	return r
}
