package adapters

import (
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"

	"crypto-viewer/src/config"
)

func NewCoins(r *resty.Client) Coins {
	return Coins{
		RestyClientAddress: r,
	}
}

type Coins struct {
	RestyClientAddress *resty.Client
}

func (c Coins) GetCoins(p params) (http.Response, error) {
	resp, err := c.RestyClientAddress.R().
		EnableTrace().
		SetQueryParams(p).
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
