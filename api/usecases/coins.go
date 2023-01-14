package usecases

import (
	"crypto-viewer/src/entities"
	"log"
)

type CoinsAdapter interface {
	GetCoins(map[string]string) (entities.CoinsData, error)
}

type ExchangeAdapter interface {
	GetExchangeRate() (entities.ExchangeRate, error)
}

func NewCoins(coinsAdapter CoinsAdapter, exchangeAdapter ExchangeAdapter) CoinsUseCase {
	return CoinsUseCase{
		coinsAdapter:    coinsAdapter,
		exchangeAdapter: exchangeAdapter,
	}
}

type CoinsUseCase struct {
	coinsAdapter    CoinsAdapter
	exchangeAdapter ExchangeAdapter
}

func (c CoinsUseCase) GetCoins(params map[string]string) (entities.CoinsData, error) {

	coinsData, err := c.coinsAdapter.GetCoins(params)
	if err != nil {
		log.Print(err)
		return entities.CoinsData{}, err
	}

	////Make exange from USD to UAH
	exchangeRate, err := c.exchangeAdapter.GetExchangeRate()
	for i := 0; i < len(coinsData.Coins); i++ {
		coinsData.Coins[i].Quote.USD.Price = coinsData.Coins[i].Quote.USD.Price * exchangeRate.Quotes.USDUAH
	}
	if err != nil {
		log.Print(err)
		return entities.CoinsData{}, err
	}

	return coinsData, nil
}
