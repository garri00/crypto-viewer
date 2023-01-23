package adapters

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"

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
		log.Print(err)
		return entities.CoinsData{}, err
	}

	//Check response code from CMC API
	if resp.StatusCode() != http.StatusOK {
		errResponse := entities.Status{}
		if err := json.Unmarshal(resp.Body(), &errResponse); err != nil {
			log.Print(err)
			log.Print("failed to unmarshal errResponse")
			return entities.CoinsData{}, err
		}
		err := fmt.Errorf("Code=%d, Message=%s", errResponse.Status.ErrorCode, errResponse.Status.ErrorMessage)
		log.Print(err)
		return entities.CoinsData{}, err
	}

	var coinsData = entities.CoinsData{}
	if err := json.Unmarshal(resp.Body(), &coinsData); err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")
		return entities.CoinsData{}, err
	}

	return coinsData, nil
}
