package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configs struct {
	CoinMarCapTokenAPI string `envconfig:"CMCTOKEN"`
	ExchangeTokenAPI   string `envconfig:"EXCHANGETOKEN"`
	LogLevel           string `envconfig:"LOGLEVEL"`
	MongoConf          MongoDBConf
	PostgreConf        PostgreDBConf
}

type MongoDBConf struct {
	Host     string `envconfig:"MGHOST" default:"localhost"`
	Port     string `envconfig:"MGPORT" default:"27017"`
	Database string `envconfig:"MGDATABASE" default:"admin"`
	Username string `envconfig:"MGUSR"`
	Password string `envconfig:"MGPAS"`
}

type PostgreDBConf struct {
	Host     string `envconfig:"PGHOST" default:"localhost"`
	Port     string `envconfig:"PGPORT" default:"5432"`
	Database string `envconfig:"PGDATABASE" default:"crypto"`
	Username string `envconfig:"PGUSR"`
	Password string `envconfig:"PGPAS"`
}

func GetConfig() (Configs, error) {
	var cfg Configs

	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("myApp", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
