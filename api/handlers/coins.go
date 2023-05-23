package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/rs/zerolog"

	"crypto-viewer/src/entities"
)

//go:generate mockgen -source=./coins.go -destination=./mock_test.go -package=handlers

type CoinsUseCase interface {
	GetCoins(params map[string]string) (entities.CoinsData, error)
}

type SaveDataUseCase interface {
	SaveCoins(coinsData entities.CoinsData) error
}

type SaveCoinsDB interface {
	SaveCoinsDB(coinsData entities.CoinsData) error
}

type CoinsHandler struct {
	coinsUseCase    CoinsUseCase
	saveDataUseCase SaveDataUseCase
	saveCoinsDB     SaveCoinsDB
	log             zerolog.Logger
}

func CoinsHendler(coinsUseCase CoinsUseCase, saveDataUseCase SaveDataUseCase, saveCoinsDB SaveCoinsDB, l zerolog.Logger) CoinsHandler {
	return CoinsHandler{
		coinsUseCase:    coinsUseCase,
		saveDataUseCase: saveDataUseCase,
		saveCoinsDB:     saveCoinsDB,
		log:             l,
	}
}

func (c CoinsHandler) CoinsResty(w http.ResponseWriter, r *http.Request) {
	queryParams := map[string]string{
		"start": r.URL.Query().Get("start"),
		"limit": r.URL.Query().Get("limit"),
	}

	c.log.Debug().Any("query params", queryParams).Msg("query params")

	if err := validateParams(queryParams); err != nil {
		c.log.Error().Err(err).Msg("wrong query params:")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("wrong query pqrams"))
		if err != nil {
			return
		}

		return
	}

	resp, err := c.coinsUseCase.GetCoins(queryParams)
	if err != nil {
		c.log.Error().Err(err).Msg("failed to create GET coins")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("failed to create GET coins"))
		if err != nil {
			return
		}

		return
	}

	if err := c.saveDataUseCase.SaveCoins(resp); err != nil {
		c.log.Error().Err(err).Msg("failed to save coins")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("failed to save coins"))
		if err != nil {
			return
		}

		return
	}

	if err := c.saveCoinsDB.SaveCoinsDB(resp); err != nil {
		c.log.Error().Err(err).Msg("failed to save coins into db")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	body, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		c.log.Error().Err(err).Msg("failed to unmarshal coinsData")

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		return
	}
}
func validateParams(params map[string]string) error {
	startValue, err := strconv.Atoi(params["start"])
	if err != nil {
		return errors.New("wrong start param")
	}
	if startValue < 1 {
		return errors.New("wrong start param < 1")
	}

	limitValue, err := strconv.Atoi(params["limit"])
	if err != nil {
		return errors.New("wrong limit param")
	}
	if limitValue < 1 || limitValue > 5000 {
		return errors.New("wrong limit param")
	}

	return nil
}
