package usecases

import (
	"crypto-viewer/src/entities"
	"log"
)

type Adapter interface {
	CoinsRestyAdapter
	ExchangeAdapter
}

type CoinsRestyAdapter interface {
	GetCoinsA(map[string]string) (entities.CoinsData, error)
}

type ExchangeAdapter interface {
	GetExchangeRateA() (entities.ExchangeRate, error)
}

func NewCoinsUC(coinsAdapter Adapter) CoinsUsecase {
	return CoinsUsecase{
		coinsA: coinsAdapter,
	}
}

type CoinsUsecase struct {
	coinsA Adapter
}

func (c CoinsUsecase) GetCoinsUC(params map[string]string) (entities.CoinsData, error) {

	coinsData, err := c.coinsA.GetCoinsA(params)
	if err != nil {
		log.Print(err)
		return coinsData, err
	}

	////Make exange from USD to UAH
	exchangeRate, err := c.coinsA.GetExchangeRateA()
	for i := 0; i < len(coinsData.Coins); i++ {
		coinsData.Coins[i].Quote.USD.Price = coinsData.Coins[i].Quote.USD.Price * exchangeRate.Quotes.USDUAH
	}

	return coinsData, nil
}
