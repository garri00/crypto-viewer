package config

import "os"

type Configs struct {
	coinMarCapTokenAPI string
	exchangeTokenAPI   string
}

func (c Configs) GetCoinMarketTokenAPI() string {
	return c.coinMarCapTokenAPI
}

func (c Configs) GetExchangeTokenAPI() string {
	return c.exchangeTokenAPI
}

var Config = Configs{
	coinMarCapTokenAPI: os.Getenv("CMCTOKEN"),
	exchangeTokenAPI:   os.Getenv("EXCHANGETOKEN"),
}
