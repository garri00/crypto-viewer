package usecases

import (
	"net/http"

	"crypto-viewer/src/entities"
)

type CoinsRestyAdapter interface {
	GetCoins() (http.Response, error)
}

func NewCoins(adapter CoinsRestyAdapter) Coins {
	return Coins{
		coinsA: adapter,
	}
}

type Coins struct {
	coinsA CoinsRestyAdapter
}

func (c Coins) GetCoins(p params) (entities.CoinsData, error) {
	p := map[string]string{
		"start": "1",
		"limit": "500",
	}

	coins, err := c.coinsA.GetCoins(p)

	return coins, nil
}

http request -> handler -> usecase -> adapter -> usecase -> handler -> hhtp response
