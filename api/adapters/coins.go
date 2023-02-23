package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
)

func NewCoins(r *resty.Client, c config.Configs) CoinsAdapter {
	return CoinsAdapter{
		restyClient: r,
		config:      c,
	}
}

type CoinsAdapter struct {
	restyClient *resty.Client
	config      config.Configs
	log         zerolog.Logger
}

func (c CoinsAdapter) GetCoins(params map[string]string) (entities.CoinsData, error) {
	// Make call to CMC API
	resp, err := c.restyClient.R().
		EnableTrace().
		SetQueryParams(params).
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", c.config.CoinMarCapTokenAPI).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

	// Check if request is OK
	if err != nil {
		err := fmt.Errorf("cant get request to CMC: %w", err)
		c.log.Error().Err(err).Msgf("")

		return entities.CoinsData{}, err
	}

	// Check response code from CMC API
	if resp.StatusCode() != http.StatusOK {
		errResponse := entities.Status{}
		if err := json.Unmarshal(resp.Body(), &errResponse); err != nil {
			err := fmt.Errorf("failed to unmarshal errResponse: %w", err)
			c.log.Error().Err(err).Msgf("")

			return entities.CoinsData{}, err
		}
		err := fmt.Errorf("code=%d, message=%s", errResponse.Body.ErrorCode, errResponse.Body.ErrorMessage)
		c.log.Error().Err(err).Msgf("")

		return entities.CoinsData{}, err
	}

	coinsData := entities.CoinsData{}
	if err := json.Unmarshal(resp.Body(), &coinsData); err != nil {
		err := fmt.Errorf("failed to unmarshal coinsData: %w", err)
		c.log.Error().Err(err).Msgf("")

		return entities.CoinsData{}, err
	}

	return coinsData, nil
}
