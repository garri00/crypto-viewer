package usecases

import (
	"crypto-viewer/src/entities"
	"fmt"
	"log"
)

//go:generate mockgen -source=./coins.go -destination=./mock_test.go -package=usecases

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
	//Get coins from CMC api
	coinsData, err := c.coinsAdapter.GetCoins(params)
	if err != nil {
		err := fmt.Errorf("cant call coins adapter: %w", err)
		log.Print(err)
		return entities.CoinsData{}, err
	}

	//Get exchange rate USD to UAH
	exchangeRate, err := c.exchangeAdapter.GetExchangeRate()
	if err != nil {
		err := fmt.Errorf("cant call exchange adapter: %w", err)
		log.Print(err)
		return entities.CoinsData{}, err
	}

	// Change coins value
	makeExchange(coinsData, exchangeRate)

	return coinsData, nil
}

// TODO: change signature
func makeExchange(coinsData entities.CoinsData, exchangeRate entities.ExchangeRate) {
	for i := 0; i < len(coinsData.Coins); i++ {
		coinsData.Coins[i].Quote.USD.Price = coinsData.Coins[i].Quote.USD.Price * exchangeRate.Quotes.USDUAH
	}
}
