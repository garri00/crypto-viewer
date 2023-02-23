package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"crypto-viewer/api"
	"crypto-viewer/src/config"
	"crypto-viewer/src/logging"
)

func main() {
	configs := config.GetConfig()

	logging.SetLogger(configs)

	log.Info().Msg("server started")

	r := api.NewRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Info().Msg("server crashed")
		return
	}
	log.Info().Msg("server stopped")
}
