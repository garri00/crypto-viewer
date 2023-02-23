package config

import "os"

type Configs struct {
	CoinMarCapTokenAPI string
	ExchangeTokenAPI   string
	LogLevel           string
}

func GetConfig() Configs {
	return сonfig
}

var сonfig = Configs{
	CoinMarCapTokenAPI: os.Getenv("CMCTOKEN"),
	ExchangeTokenAPI:   os.Getenv("EXCHANGETOKEN"),
	LogLevel:           os.Getenv("LOGLEVEL"),
}
