package handlers

import (
	"crypto-viewer/api/handlers/usecases"
	"net/http"
)

// CoinsUsecase represent Coins use-case layer

type CoinsHendlerContract struct {
	coinsUC usecases.CoinsUsecase
}

func NewCoinsHendler(usecase usecases.CoinsUsecase) CoinsHendlerContract {
	return CoinsHendlerContract{
		coinsUC: usecase,
	}
}

type CoinsHendlerH interface {
	CoinsResty(w http.ResponseWriter, r *http.Request)
}
