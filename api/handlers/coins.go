package handlers

import (
	"log"
	"net/http"

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

	// TODO: validate params

	resp, err := c.coinsUseCase.GetCoins(queryParams)

	// TODO: make err wrap
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET coins"))
		return
	}

	file, err := c.saveDataUseCase.SaveCoins(resp)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to save coins"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(file)

}
