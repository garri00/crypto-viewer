package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"crypto-viewer/src/entities"
)

//go:generate mockgen -source=./coins.go -destination=./mock_test.go -package=usecases

type CoinsUseCase interface {
	GetCoins(params map[string]string) (entities.CoinsData, error)
}

type SaveDataUseCase interface {
	SaveCoins(coinsData entities.CoinsData) ([]byte, error)
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
		err := fmt.Errorf("wrong query params %w", err)
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET coins"))
		return
	}

	resp, err := c.coinsUseCase.GetCoins(queryParams)
	if err != nil {
		err := fmt.Errorf("failed to create GET coins %w", err)
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET coins"))
		return
	}

	file, err := c.saveDataUseCase.SaveCoins(resp)
	if err != nil {
		err := fmt.Errorf("failed to save coins %w", err)
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to save coins"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(file)

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
