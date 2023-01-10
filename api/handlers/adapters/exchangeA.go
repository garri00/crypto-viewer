package adapters

import (
	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"log"
)

func (c CoinsAdapter) GetExchangeRateA() (entities.ExchangeRate, error) {

	exchangeRate := entities.ExchangeRate{}
	resp, err := c.RestyClientAddress.R().
		SetQueryParams(map[string]string{
			"source":     "USD",
			"currencies": "UAH",
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("apikey", config.GetExchangeTokenAPI()).
		Get("https://api.apilayer.com/currency_data/live")

	if err != nil {
		log.Print(err)
		return exchangeRate, err
	}

	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
		log.Print(err)
		log.Print("failed to unmarshal exchangeRateData")
		fmt.Println(resp.Body())
		return exchangeRate, err
	}

	if exchangeRate.Success != true {
		log.Print(err)
		log.Print("cant get exchange rate data")

		return exchangeRate, err
	}
	return exchangeRate, nil
}
