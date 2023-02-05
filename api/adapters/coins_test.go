package adapters

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var params = map[string]string{
	"start": "1",
	"limit": "4",
}

var requestPath = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"

var okResponse = entities.CoinsData{

	Coins: []entities.Coin{
		{
			Id:     1,
			Name:   "BitCoin",
			Symbol: "BTC",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 20,
				},
			},
		},
		{
			Id:     300,
			Name:   "Etherym",
			Symbol: "ETH",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 10,
				},
			},
		},
		{
			Id:     341,
			Name:   "BNB",
			Symbol: "BNB",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 11255.190172461225,
				},
			},
		},
		{
			Id:     1233,
			Name:   "Tether",
			Symbol: "USDT",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 36.77161508245386,
				},
			},
		},
	},
}

var errMessage = entities.Status{
	Body: entities.Body{
		ErrorCode:    400,
		ErrorMessage: "Invalid value for \"id\"",
	},
}

func TestCoinsAdapter_GetCoins(t *testing.T) {
	restyClient := resty.New()

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	tests := map[string]struct {
		restyClient *resty.Client
		responder   func()
		exp         entities.CoinsData
		expErr      error
	}{
		"succes": {
			restyClient: restyClient,
			responder: func() {
				jsonBytes, _ := json.Marshal(okResponse)
				responder := httpmock.NewStringResponder(200, string(jsonBytes))
				httpmock.RegisterResponderWithQuery("GET", requestPath, "start=1&limit=4", responder)
			},
			exp:    okResponse,
			expErr: nil,
		},

		"bad responce from CMC api": {
			restyClient: restyClient,
			responder: func() {
				jsonBytes, _ := json.Marshal(errMessage)
				responder := httpmock.NewStringResponder(400, string(jsonBytes))
				httpmock.RegisterResponderWithQuery("GET", requestPath, "start=1&limit=4", responder)
			},
			exp:    entities.CoinsData{},
			expErr: errors.New("Code=400, Message=Invalid value for \"id\""),
		},

		"bad responce from CMC api: failed to unmarshal errResponse": {
			restyClient: restyClient,
			responder: func() {
				responder := httpmock.NewStringResponder(500, "FAIL")
				httpmock.RegisterResponderWithQuery("GET", requestPath, "start=1&limit=4", responder)
			},
			exp:    entities.CoinsData{},
			expErr: errors.New("failed to unmarshal errResponse: invalid character 'F' looking for beginning of value"),
		},

		"bad responce from CMC api: failed to unmarshal coinsData": {
			restyClient: restyClient,
			responder: func() {
				responder := httpmock.NewStringResponder(200, "COINS")
				httpmock.RegisterResponderWithQuery("GET", requestPath, "start=1&limit=4", responder)
			},
			exp:    entities.CoinsData{},
			expErr: errors.New("failed to unmarshal coinsData: invalid character 'C' looking for beginning of value"),
		},

		"bad can`t get unswer form api": {
			restyClient: restyClient,
			responder: func() {
				respBody := errors.New("Get \"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=4&start=1\": dial tcp: lookup pro-api.coinmarketcap.c\nom: no such host")
				responder := httpmock.NewErrorResponder(respBody)
				httpmock.RegisterResponderWithQuery("GET", requestPath, "start=1&limit=4", responder)
			},
			exp:    entities.CoinsData{},
			expErr: errors.New("cant get request to CMC: Get \"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=4&start=1\": Get \"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=4&start=1\": dial tcp: lookup pro-api.coinmarketcap.c\nom: no such host"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := CoinsAdapter{
				restyClient: tt.restyClient,
			}

			tt.responder()
			got, err := c.GetCoins(params)
			assert.Equal(t, tt.exp, got)
			if tt.expErr != nil {
				assert.EqualError(t, err, tt.expErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
