package adapters

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"

	"crypto-viewer/src/config"
)

func NewCoinsAdapter(r *resty.Client) CoinsAdapter {
	return CoinsAdapter{
		RestyClientAddress: r,
	}
}

type CoinsAdapter struct {
	RestyClientAddress *resty.Client
}

func (c CoinsAdapter) GetCoinsA(params map[string]string) (entities.CoinsData, error) {

	//Make call to CMC API
	resp, err := c.RestyClientAddress.R().
		EnableTrace().
		SetQueryParams(params).
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", config.GetConfigTokenAPI()).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

	fmt.Println(resp.Request.URL)
	//Check if request is OK
	var coinsData = entities.CoinsData{}
	if err != nil {
		log.Print(err)
		return coinsData, err
	}
	//Check response code from CMC API
	var errResponse = entities.Status{}
	if resp.StatusCode() != http.StatusOK {
		if err := json.Unmarshal(resp.Body(), &errResponse); err != nil {
			log.Print(err)
			log.Print("failed to unmarshal errResponse")
			return coinsData, err
		}
		log.Printf("Code=%d, Message=%s", errResponse.Status.ErrorCode, errResponse.Status.ErrorMessage)
		return coinsData, err
	}

	if err := json.Unmarshal(resp.Body(), &coinsData); err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")
		return coinsData, err
	}

	//Write coinsData to file
	file, err := json.MarshalIndent(coinsData, "", " ")
	if err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")
		return coinsData, err
	}
	ioutil.WriteFile("src/pkg/coinslist.json", file, 0644)

	return coinsData, nil
}
