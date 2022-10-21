package tests

import (
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"crypto-viewer/api"
	"crypto-viewer/src/entities"
	"github.com/gavv/httpexpect/v2"
)

func Test_CoinsRestyHandlerHTTPEXP(t *testing.T) {

	var okResponse = entities.Data{}
	jsonFile, err := os.Open("coinslist_test.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &okResponse)

	// run server using httptest

	router := api.NewRouter()
	httpClient := &http.Client{
		Transport: httpexpect.NewBinder(router),
		Jar:       httpexpect.NewJar(),
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Client:   httpClient,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	//

	// is it working?
	e.GET("/coins").
		WithQuery("start", "1").
		WithQuery("limit", "3").
		Expect().Status(http.StatusOK)

	e.GET("").
		Expect().
		Status(http.StatusOK)

	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", httpmock.NewStringResponder(200, j))

}

const j = `{
 "data": [
  {
   "id": 1,
   "name": "Bitcoin",
   "symbol": "BTC",
   "slug": "bitcoin",
   "num_market_pairs": 0001,
   "date_added": "2013-04-28T00:00:00Z",
   "max_supply": 21000000,
   "circulating_supply": 19169075,
   "total_supply": 19169075,
   "quote": {
    "USD": {
     "price": 20000.26633700906,
     "volume_24h": 34057182267.847286,
     "volume_change_24h": 16.555,
     "percent_change_1h": -0.47666159,
     "percent_change_24h": 2.94124922,
     "percent_change_7d": 4.93977254,
     "percent_change_30d": 0.4471352,
     "percent_change_60d": -13.10799526,
     "percent_change_90d": -1.56798353,
     "market_cap": 383386605434.1019,
     "market_cap_dominance": 39.9464,
     "fully_diluted_market_cap": 420005593077.19,
     "last_updated": "2022-10-04T17:04:00Z"
    }
   }
  },
  {
   "id": 1027,
   "name": "Ethereum",
   "symbol": "ETH",
   "slug": "ethereum",
   "num_market_pairs": 6121,
   "date_added": "2015-08-07T00:00:00Z",
   "max_supply": null,
   "circulating_supply": 122650819.499,
   "total_supply": 122650819.499,
   "quote": {
    "USD": {
     "price": 1347.8692249214917,
     "volume_24h": 9870895591.756325,
     "volume_change_24h": -4.808,
     "percent_change_1h": -0.43006632,
     "percent_change_24h": 2.05465002,
     "percent_change_7d": 1.66856006,
     "percent_change_30d": -14.088573,
     "percent_change_60d": -19.35345636,
     "percent_change_90d": 17.74305407,
     "market_cap": 165317265014.1029,
     "market_cap_dominance": 17.225,
     "fully_diluted_market_cap": 165317265014.1,
     "last_updated": "2022-10-04T17:04:00Z"
    }
   }
  },
  {
   "id": 825,
   "name": "Tether",
   "symbol": "USDT",
   "slug": "tether",
   "num_market_pairs": 40475,
   "date_added": "2015-02-25T00:00:00Z",
   "max_supply": null,
   "circulating_supply": 67949574445.85782,
   "total_supply": 70155449906.09836,
   "quote": {
    "USD": {
     "price": 1.000221198547823,
     "volume_24h": 42644483506.99466,
     "volume_change_24h": 8.6216,
     "percent_change_1h": 0.000744,
     "percent_change_24h": 0.01765955,
     "percent_change_7d": 0.019885,
     "percent_change_30d": 0.01800994,
     "percent_change_60d": 0.01539732,
     "percent_change_90d": 0.10456315,
     "market_cap": 67964604793.05044,
     "market_cap_dominance": 7.0813,
     "fully_diluted_market_cap": 70170968189.74,
     "last_updated": "2022-10-04T17:04:00Z"
    }
   }
  }
 ]
}`
