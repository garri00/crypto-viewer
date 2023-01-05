package handlers

import (
	"net/http"

	"crypto-viewer/src/entities"
)

// CoinsUsecase represent Coins use-case layer

//type CoinsUsecase interface {
//	GetCoins(r *http.Request) error
//}

type CoinsRestyUseCase interface {
	GetCoins(http.Response) (entities.CoinsData, error)
}
