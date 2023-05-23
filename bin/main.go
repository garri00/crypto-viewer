package main

import (
	"net/http"
	"time"

	"crypto-viewer/api"
	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/config"
)

func main() {
	configs, err := config.GetConfig()
	if err != nil {
		logger.Log.Fatal().Msg("failed to load .env")
	}

	logger.SetLogLevel(configs)

	logger.Log.Info().Msg("server started")

	r := api.NewRouter(configs)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Log.Info().Msg("server crashed")
		panic(err)
	}

}
