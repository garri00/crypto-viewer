package adapters

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rs/zerolog"

	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"

	"github.com/go-resty/resty/v2"
)

func NewExchange(r *resty.Client, c config.Configs) ExchangeAdapter {
	return ExchangeAdapter{
		restyClient: r,
		config:      c,
	}
}

type ExchangeAdapter struct {
	restyClient *resty.Client
	config      config.Configs
	log         zerolog.Logger
}

func (c ExchangeAdapter) GetExchangeRate() (entities.ExchangeRate, error) {
	resp, err := c.restyClient.R().
		SetQueryParams(map[string]string{
			"source":     "USD",
			"currencies": "UAH",
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("apikey", config.GetConfig().ExchangeTokenAPI).
		Get("https://api.apilayer.com/currency_data/live")

	if err != nil {
		err := fmt.Errorf("cant call exchange api: %w", err)
		c.log.Error().Err(err).Msgf("")

		return entities.ExchangeRate{}, err
	}

	exchangeRate := entities.ExchangeRate{}
	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
		err := fmt.Errorf("failed to unmarshal exchangeRateData: %w", err)
		c.log.Error().Err(err).Msgf("")

		return entities.ExchangeRate{}, err
	}

	if !exchangeRate.Success {
		err := errors.New("can`t get exchange rate data")
		c.log.Error().Err(err).Msgf("")

		return entities.ExchangeRate{}, err
	}

	return exchangeRate, nil
}
