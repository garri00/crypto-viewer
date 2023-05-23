package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"

	"crypto-viewer/api/adapters"
	"crypto-viewer/api/adapters/db/mongo"
	"crypto-viewer/api/adapters/db/postgres"
	"crypto-viewer/api/handlers"
	"crypto-viewer/api/usecases"
	"crypto-viewer/pkg/clients/mongodb"
	"crypto-viewer/pkg/clients/posgresql"
	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/config"
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

	ctx := context.Background()

	mongoClient, err := mongodb.NewClient(ctx, configs.MongoConf)
	if err != nil {
		logger.Log.Err(err).Msg("failed to create mongo client")
		return nil
	}
	coinsCollection := mongo.NewStorageMG("crypto", mongoClient, logger.Log)

	postgresClient, err := posgresql.NewClient(ctx, configs.PostgreConf)
	if err != nil {
		logger.Log.Err(err).Msg("failed to create postgres")
	}

	_ = postgres.NewStoragePG(postgresClient, logger.Log)

	restyClient := resty.New()
	coinsAdapter := adapters.NewCoins(restyClient, configs, logger.Log)
	exchangeAdapter := adapters.NewExchange(restyClient, configs, logger.Log)

	coinsUseCase := usecases.NewCoins(coinsAdapter, exchangeAdapter, logger.Log)
	saveDataUseCase := usecases.NewSaveData(logger.Log)
	saveCoinsDBUsecase := usecases.NewSaveCoinsDB(coinsCollection, logger.Log)

	c := handlers.CoinsHendler(coinsUseCase, saveDataUseCase, saveCoinsDBUsecase, logger.Log)

	r.Get("/home", handlers.HomeHandler)
	r.Get("/coins", c.CoinsResty)

	return r
}
