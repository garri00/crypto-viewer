package usecases

import (
	"crypto-viewer/src/entities"
	"log"
)

type CoinsRestyAdapter interface {
	GetCoinsA(map[string]string) (entities.CoinsData, error)
}

type ExchangeAdapter interface {
	GetExchangeRateA() (entities.ExchangeRate, error)
}

func NewCoinsUC(coinsAdapter CoinsRestyAdapter, ExchangeAdapter ExchangeAdapter) CoinsUsecase {
	return CoinsUsecase{
		coinsA:          coinsAdapter,
		exchangeAdapter: ExchangeAdapter,
	}
}

type CoinsUsecase struct {
	coinsA          CoinsRestyAdapter
	exchangeAdapter ExchangeAdapter
}

func (c CoinsUsecase) GetCoinsUC(params map[string]string) (entities.CoinsData, error) {

	coinsData, err := c.coinsA.GetCoinsA(params)
	if err != nil {
		log.Print(err)
		return coinsData, err
	}

	////Make exange from USD to UAH
	exchangeRate, err := c.exchangeAdapter.GetExchangeRateA()
	for i := 0; i < len(coinsData.Coins); i++ {
		coinsData.Coins[i].Quote.USD.Price = coinsData.Coins[i].Quote.USD.Price * exchangeRate.Quotes.USDUAH
	}

	return coinsData, nil
}

//http request -> handler -> usecase -> adapter -> usecase -> handler -> hhtp response
