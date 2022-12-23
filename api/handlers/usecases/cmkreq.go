package usecases

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type CoinsRestyUseCase interface {
	GetCoins(r resty.Client) (http.Response, error)
}

type Coins struct {
	coinsUC CoinsRestyUseCase
}

func (c Coins) GetCoins(r resty.Client) error {

	coins, err := c.coinsUC.GetCoins(r)
	if err != nil {
		return err
	}
	return coins

}
