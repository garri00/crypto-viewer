package usecases

import (
	"fmt"

	"github.com/rs/zerolog"

	"crypto-viewer/src/entities"
)

//go:generate mockgen -source=./coins.go -destination=./mock_test.go -package=usecases

type CoinsAdapter interface {
	GetCoins(map[string]string) (entities.CoinsData, error)
}

type ExchangeAdapter interface {
	GetExchangeRate() (entities.ExchangeRate, error)
}

func NewCoins(coinsAdapter CoinsAdapter, exchangeAdapter ExchangeAdapter, l zerolog.Logger) CoinsUseCase {
	return CoinsUseCase{
		coinsAdapter:    coinsAdapter,
		exchangeAdapter: exchangeAdapter,
		log:             l,
	}
}

type CoinsUseCase struct {
	coinsAdapter    CoinsAdapter
	exchangeAdapter ExchangeAdapter
	log             zerolog.Logger
}

func (c CoinsUseCase) GetCoins(params map[string]string) (entities.CoinsData, error) {
	// Get coins from CMC api
	coinsData, err := c.coinsAdapter.GetCoins(params)
	if err != nil {
		err := fmt.Errorf("cant call coins adapter: %w", err)
		c.log.Err(err).Msg("")

		return entities.CoinsData{}, err
	}

	// Get exchange rate USD to UAH
	exchangeRate, err := c.exchangeAdapter.GetExchangeRate()
	if err != nil {
		err := fmt.Errorf("cant call exchange adapter: %w", err)
		c.log.Err(err).Msg("")

		return entities.CoinsData{}, err
	}

	// Change coins value
	coinsDataUAH := makeExchange(coinsData, exchangeRate)

	return coinsDataUAH, nil
}

func makeExchange(coinsData entities.CoinsData, exchangeRate entities.ExchangeRate) entities.CoinsData {
	for i := 0; i < len(coinsData.Coins); i++ {
		coinsData.Coins[i].Quote.USD.Price *= exchangeRate.Quotes.USDUAH
	}

	return coinsData
}
