package adapters

import (
	"crypto-viewer/src/config"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type CoinsRestyAdapter interface {
	GetCoins() (http.Response, error)
}

type Coins struct {
	coinsA             CoinsRestyAdapter
	RestyClientAddress *resty.Client
}

func (c Coins) GetCoins() (http.Response, error) {

	restyClient := c.RestyClientAddress

	resp, err := restyClient.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"start": "1",
			"limit": "500",
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", config.TokenAPI).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

	coins := *resp.RawResponse
	//fmt.Println(resp.Request.URL)

	if err != nil {
		log.Print(err)
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}
	return coins, nil

}
