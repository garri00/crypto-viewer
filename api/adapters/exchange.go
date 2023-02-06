package adapters

import (
	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

func NewExchange(r *resty.Client) ExchangeAdapter {
	return ExchangeAdapter{
		restyClient: r,
	}
}

type ExchangeAdapter struct {
	restyClient *resty.Client
}

func (c ExchangeAdapter) GetExchangeRate() (entities.ExchangeRate, error) {

	resp, err := c.restyClient.R().
		SetQueryParams(map[string]string{
			"source":     "USD",
			"currencies": "UAH",
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("apikey", config.Config.GetExchangeTokenAPI()).
		Get("https://api.apilayer.com/currency_data/live")

	if err != nil {
		err := fmt.Errorf("cant call exchange api: %w", err)
		log.Print(err)
		return entities.ExchangeRate{}, err
	}

	exchangeRate := entities.ExchangeRate{}
	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
		err := fmt.Errorf("failed to unmarshal exchangeRateData: %w", err)
		log.Print(err)
		return entities.ExchangeRate{}, err
	}

	if exchangeRate.Success != true {
		err := errors.New("can`t get exchange rate data")
		log.Print(err)
		return entities.ExchangeRate{}, err
	}

	return exchangeRate, nil
}
