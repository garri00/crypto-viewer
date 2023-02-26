package main

import (
	"net/http"
	"time"

	"crypto-viewer/api"
	"crypto-viewer/src/config"
	"crypto-viewer/src/logger"
)

func main() {
	configs := config.GetConfig()

	logger.SetLogLevel(configs)

	logger.Log.Info().Msg("server started")

	r := api.NewRouter()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Log.Info().Msg("server crashed")
		panic(err)
	}
}
