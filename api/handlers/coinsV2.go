package handlers

import (
	"net/http"

	"crypto-viewer/src/entities"
)

type CoinsRestyUseCase interface {
	GetCoins(http.Response) (entities.CoinsData, error)
}

func (c RestyClient) CoinsRestyV2(w http.ResponseWriter, r *http.Request) {

}
