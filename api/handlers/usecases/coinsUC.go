package usecases

import (
	"crypto-viewer/src/entities"
	"log"
)

type CoinsRestyAdapter interface {
	GetCoinsA(map[string]string) (entities.CoinsData, error)
}

func NewCoins(adapter CoinsRestyAdapter) CoinsUC {
	return CoinsUC{
		coinsA: adapter,
	}
}

type CoinsUC struct {
	coinsA CoinsRestyAdapter
}

func (c CoinsUC) GetCoinsUC(params map[string]string) (entities.CoinsData, error) {

	coinsData, err := c.coinsA.GetCoinsA(params)
	if err != nil {
		log.Print(err)
		return coinsData, err
	}
	return coinsData, nil
}

//http request -> handler -> usecase -> adapter -> usecase -> handler -> hhtp response
