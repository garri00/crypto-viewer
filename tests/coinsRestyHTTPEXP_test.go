package tests

import (
	"crypto-viewer/api"
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_CoinsRestyHandlerHTTPEXP(t *testing.T) {

	var okResponse = entities.Data{}
	jsonFile, err := os.Open("coinslist_test.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &okResponse)

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

	//	var handler fasthttp.Client = httpClient

	// Тут я намагаюся мокнути та змінити транспорт на роутер але він починає використовувати транспорт з httpmock і ендпоінти не відпрацьовують
	httpmock.ActivateNonDefault(httpClient)
	//httpmock.Activate()

	defer httpmock.DeactivateAndReset()
	//httpmock.InitialTransport = httpClient.Transport
	responder := httpmock.NewStringResponder(200, j)
	httpmock.RegisterResponder("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=3&start=1", responder)

	//
	////httpmock.RegisterResponder("Get","/",)
	//
	//httpmock.RegisterResponder(
	//	"GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest",
	//	httpmock.NewStringResponder(200, j))

	//httpmock.Reset()
	// is it working?

	//ContentType("application", "json")

	//server := httptest.NewServer(router)
	//defer server.Close()
	//e := httpexpect.New(t, server.URL)

	e.GET("/coins").
		WithQuery("start", "1").
		WithQuery("limit", "3").
		Expect().
		Status(http.StatusOK).JSON()

	e.GET("").
		Expect().
		Status(http.StatusOK)

}

const j = `{
"data": [
{
"id": 1,
"name": "BBBB",
"symbol": "BTC",
"slug": "bitcoin",
"cmc_rank": 5,
"num_market_pairs": 500,
"circulating_supply": 16950100,
"total_supply": 16950100,
"max_supply": 21000000,
"last_updated": "2018-06-02T22:51:28.209Z",
"date_added": "2013-04-28T00:00:00.000Z",
"tags": [
"mineable"
],
"platform": null,
"self_reported_circulating_supply": null,
"self_reported_market_cap": null,
"quote": {
"USD": {
"price": 9283.92,
"volume_24h": 7155680000,
"volume_change_24h": -0.152774,
"percent_change_1h": -0.152774,
"percent_change_24h": 0.518894,
"percent_change_7d": 0.986573,
"market_cap": 852164659250.2758,
"market_cap_dominance": 51,
"fully_diluted_market_cap": 952835089431.14,
"last_updated": "2018-08-09T22:53:32.000Z"
},
"BTC": {
"price": 1,
"volume_24h": 772012,
"volume_change_24h": 0,
"percent_change_1h": 0,
"percent_change_24h": 0,
"percent_change_7d": 0,
"market_cap": 17024600,
"market_cap_dominance": 12,
"fully_diluted_market_cap": 952835089431.14,
"last_updated": "2018-08-09T22:53:32.000Z"
}
}
},
{
"id": 1027,
"name": "Ethereum",
"symbol": "ETH",
"slug": "ethereum",
"num_market_pairs": 6360,
"circulating_supply": 16950100,
"total_supply": 16950100,
"max_supply": 21000000,
"last_updated": "2018-06-02T22:51:28.209Z",
"date_added": "2013-04-28T00:00:00.000Z",
"tags": [
"mineable"
],
"platform": null,
"quote": {
"USD": {
"price": 1283.92,
"volume_24h": 7155680000,
"volume_change_24h": -0.152774,
"percent_change_1h": -0.152774,
"percent_change_24h": 0.518894,
"percent_change_7d": 0.986573,
"market_cap": 158055024432,
"market_cap_dominance": 51,
"fully_diluted_market_cap": 952835089431.14,
"last_updated": "2018-08-09T22:53:32.000Z"
},
"ETH": {
"price": 1,
"volume_24h": 772012,
"volume_change_24h": -0.152774,
"percent_change_1h": 0,
"percent_change_24h": 0,
"percent_change_7d": 0,
"market_cap": 17024600,
"market_cap_dominance": 12,
"fully_diluted_market_cap": 952835089431.14,
"last_updated": "2018-08-09T22:53:32.000Z"
}
}
}
],
"status": {
"timestamp": "2018-06-02T22:51:28.209Z",
"error_code": 0,
"error_message": "",
"elapsed": 10,
"credit_count": 1
}
}`
