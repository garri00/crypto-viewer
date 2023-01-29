package adapters

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"

	"crypto-viewer/src/config"
)

func NewCoins(r *resty.Client) CoinsAdapter {
	return CoinsAdapter{
		restyClient: r,
	}
}

type CoinsAdapter struct {
	restyClient *resty.Client
}

type Query struct {
	start string
	limit string
}

func (c CoinsAdapter) GetCoins(params map[string]string) (entities.CoinsData, error) {
	//Make call to CMC API
	resp, err := c.restyClient.R().
		EnableTrace().
		SetQueryParams(params).
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", config.GetConfigTokenAPI()).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

	//Check if request is OK
	if err != nil {
		err := fmt.Errorf("cant get request to CMC: %w", err)
		log.Print(err)
		return entities.CoinsData{}, err
	}

	//Check response code from CMC API
	if resp.StatusCode() != http.StatusOK {
		errResponse := entities.Status{}
		if err := json.Unmarshal(resp.Body(), &errResponse); err != nil {
			err := fmt.Errorf("failed to unmarshal errResponse: %w", err)
			log.Print(err)
			return entities.CoinsData{}, err
		}
		err := fmt.Errorf("Code=%d, Message=%s", errResponse.Body.ErrorCode, errResponse.Body.ErrorMessage)
		log.Print(err)
		return entities.CoinsData{}, err
	}

	var coinsData = entities.CoinsData{}
	if err := json.Unmarshal(resp.Body(), &coinsData); err != nil {
		err := fmt.Errorf("failed to unmarshal coinsData: %w", err)
		log.Print(err)
		return entities.CoinsData{}, err
	}

	return coinsData, nil
}
