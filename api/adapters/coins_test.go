package adapters

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var testCoins = entities.CoinsData{
	Coins: nil,
}

func TestCoinsAdapter_GetCoins(t *testing.T) {
	restyClient := resty.New()

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	var okResponse = entities.CoinsData{}
	jsonFile, err := os.Open("test_coinsData.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &okResponse)

	tests := map[string]struct {
		restyClient *resty.Client
		params      map[string]string
		responder   func()
		want        entities.CoinsData
		wantErr     bool
	}{
		"succes": {
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			responder: func() {
				responder := httpmock.NewStringResponder(200, j)
				httpmock.RegisterResponderWithQuery("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", "start=1&limit=4", responder)
			},
			want:    okResponse,
			wantErr: false,
		},

		"bad responce from CMC api": {
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			responder: func() {
				responder := httpmock.NewStringResponder(400, "{\n\"status\": {\n\"timestamp\": \"2018-06-02T22:51:28.209Z\",\n\"error_code\": 400,\n\"error_message\": \"Invalid value for \\\"id\\\"\",\n\"elapsed\": 10,\n\"credit_count\": 0\n}\n}")
				httpmock.RegisterResponderWithQuery("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", "start=1&limit=4", responder)
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},

		"bad responce from CMC api: failed to unmarshal errResponse": {
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			responder: func() {
				responder := httpmock.NewStringResponder(500, "FAIL")
				httpmock.RegisterResponderWithQuery("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", "start=1&limit=4", responder)
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},

		"bad responce from CMC api: failed to unmarshal coinsData": {
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			responder: func() {
				responder := httpmock.NewStringResponder(200, "COINS")
				httpmock.RegisterResponderWithQuery("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", "start=1&limit=4", responder)
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},

		"bad can`t get unswer form api": {
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			responder: func() {
				responder := httpmock.NewErrorResponder(errors.New("Get \"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=4&start=1\": dial tcp: lookup pro-api.coinmarketcap.c\nom: no such host"))
				httpmock.RegisterResponderWithQuery("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", "start=1&limit=4", responder)
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},
	}

	//TODO: check errors
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := CoinsAdapter{
				restyClient: tt.restyClient,
			}

			//assert.Equal(t, tt.exp, got)
			//if tt.expErr != nil {
			//	assert.EqualError(t, err, tt.expErr.Error())
			//} else {
			//	assert.NoError(t, err)
			//}
			tt.responder()
			got, err := c.GetCoins(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// how to edit this params
const j = "{ \"data\": [ { \"id\": 1, \"name\": \"Bitcoin\", \"symbol\": \"BTC\", \"slug\": \"bitcoin\", \"num_market_pairs\": 9942, \"date_added\": \"2013-04-28T00:00:00Z\", \"max_supply\": 21000000, \"circulating_supply\": 19275262, \"total_supply\": 19275262, \"quote\": { \"USD\": { \"price\": 846395.6636492553, \"volume_24h\": 17967799672.783348, \"volume_change_24h\": -23.7899, \"percent_change_1h\": 0.01453089, \"percent_change_24h\": -0.80224526, \"percent_change_7d\": -0.75545004, \"percent_change_30d\": 38.54452601, \"percent_change_60d\": 40.45785973, \"percent_change_90d\": 11.42421168, \"market_cap\": 443774362030.6802, \"market_cap_dominance\": 42.2975, \"fully_diluted_market_cap\": 483483005452.5, \"last_updated\": \"2023-01-28T17:56:00Z\" } } }, { \"id\": 1027, \"name\": \"Ethereum\", \"symbol\": \"ETH\", \"slug\": \"ethereum\", \"num_market_pairs\": 6384, \"date_added\": \"2015-08-07T00:00:00Z\", \"max_supply\": null, \"circulating_supply\": 122373866.2178, \"total_supply\": 122373866.2178, \"quote\": { \"USD\": { \"price\": 57953.662323160825, \"volume_24h\": 6322905509.357651, \"volume_change_24h\": -18.4193, \"percent_change_1h\": -0.09645666, \"percent_change_24h\": -1.72831204, \"percent_change_7d\": -4.69472435, \"percent_change_30d\": 31.5633756, \"percent_change_60d\": 30.42020094, \"percent_change_90d\": -0.64004454, \"market_cap\": 192911472410.17773, \"market_cap_dominance\": 18.387, \"fully_diluted_market_cap\": 192911472410.18, \"last_updated\": \"2023-01-28T17:56:00Z\" } } }, { \"id\": 825, \"name\": \"Tether\", \"symbol\": \"USDT\", \"slug\": \"tether\", \"num_market_pairs\": 46485, \"date_added\": \"2015-02-25T00:00:00Z\", \"max_supply\": null, \"circulating_supply\": 67504444776.48582, \"total_supply\": 73141766321.23428, \"quote\": { \"USD\": { \"price\": 36.77161508245386, \"volume_24h\": 27621374314.834503, \"volume_change_24h\": -15.0837, \"percent_change_1h\": -0.00000688, \"percent_change_24h\": 0.00784163, \"percent_change_7d\": -0.0010732, \"percent_change_30d\": 0.04872767, \"percent_change_60d\": 0.05016228, \"percent_change_90d\": 0.00064054, \"market_cap\": 67520175684.989456, \"market_cap_dominance\": 6.4356, \"fully_diluted_market_cap\": 73158810923.82, \"last_updated\": \"2023-01-28T17:56:00Z\" } } }, { \"id\": 1839, \"name\": \"BNB\", \"symbol\": \"BNB\", \"slug\": \"bnb\", \"num_market_pairs\": 1170, \"date_added\": \"2017-07-25T00:00:00Z\", \"max_supply\": 200000000, \"circulating_supply\": 157902256.04814193, \"total_supply\": 159979963.59042934, \"quote\": { \"USD\": { \"price\": 11255.190172461225, \"volume_24h\": 476339596.6421866, \"volume_change_24h\": 1.5379, \"percent_change_1h\": -0.22125271, \"percent_change_24h\": -0.63135193, \"percent_change_7d\": 0.77042459, \"percent_change_30d\": 24.5228616, \"percent_change_60d\": 3.49624213, \"percent_change_90d\": -2.85173237, \"market_cap\": 48342561815.94364, \"market_cap_dominance\": 4.6077, \"fully_diluted_market_cap\": 61230995713.2, \"last_updated\": \"2023-01-28T17:56:00Z\" } } } ] }"
