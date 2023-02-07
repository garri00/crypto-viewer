package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"crypto-viewer/src/entities"
)

//go:generate mockgen -source=./coins.go -destination=./mock_test.go -package=handlers

type CoinsUseCase interface {
	GetCoins(params map[string]string) (entities.CoinsData, error)
}

type SaveDataUseCase interface {
	SaveCoins(coinsData entities.CoinsData) error
}

type CoinsHandler struct {
	coinsUseCase    CoinsUseCase
	saveDataUseCase SaveDataUseCase
}

func CoinsHendler(coinsUseCase CoinsUseCase, saveDataUseCase SaveDataUseCase) CoinsHandler {
	return CoinsHandler{
		coinsUseCase:    coinsUseCase,
		saveDataUseCase: saveDataUseCase,
	}
}

func (c CoinsHandler) CoinsResty(w http.ResponseWriter, r *http.Request) {
	// Get params for Getcoins
	queryParams := map[string]string{
		"start": r.URL.Query().Get("start"),
		"limit": r.URL.Query().Get("limit"),
	}

	if err := validateParams(queryParams); err != nil {
		log.Print(fmt.Errorf("wrong query params: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong query pqrams"))
		return
	}

	resp, err := c.coinsUseCase.GetCoins(queryParams)
	if err != nil {
		log.Print(fmt.Errorf("failed to create GET coins: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET coins"))
		return
	}

	if err := c.saveDataUseCase.SaveCoins(resp); err != nil {
		log.Print(fmt.Errorf("failed to save coins: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to save coins"))
		return
	}

	body, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

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
	if limitValue < 1 && limitValue > 5000 {
		return errors.New("wrong limit param")
	}

	return nil
}
