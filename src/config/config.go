package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	CoinMarCapTokenAPI string
	ExchangeTokenAPI   string
	LogLevel           string
	PosgrePass         string
	MongoPass          string
	MongoConf          MongoDBConf
	PostgreConf        PostgreDBConf
}

type MongoDBConf struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type PostgreDBConf struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func GetConfig() (Configs, error) {
	err := godotenv.Load()
	if err != nil {
		return Configs{}, err
	}

	var сonfig = Configs{
		CoinMarCapTokenAPI: os.Getenv("CMCTOKEN"),
		ExchangeTokenAPI:   os.Getenv("EXCHANGETOKEN"),
		LogLevel:           os.Getenv("LOGLEVEL"),

		MongoConf: MongoDBConf{
			Host:     "localhost",
			Port:     "5432",
			Database: "crypto",
			Username: os.Getenv("MGUSR"),
			Password: os.Getenv("MGPAS"),
		},
		PostgreConf: PostgreDBConf{
			Host:     "locachost",
			Port:     "27017",
			Username: os.Getenv("PGUSR"),
			Password: os.Getenv("PGPAS"),
		},
	}

	return сonfig, nil
}
