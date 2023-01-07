package adapters

import (
	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

func NewExchangeAdapter(r *resty.Client) CoinsAdapter {
	return CoinsAdapter{
		RestyClientAddress: r,
	}
}

//	type ExchangeAdapter struct {
//		RestyClientAddress *resty.Client
//	}
func (c CoinsAdapter) GetExchangeRateA() (entities.ExchangeRate, error) {
	//
	//url := "https://api.apilayer.com/currency_data/live?source=USD&currencies=UAH"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("apikey", "rZKarI25VjvSBchH2Lx6TGKiBUYTqWiE")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//res, err := client.Do(req)
	//body, err := ioutil.ReadAll(res.Body)

	exchangeRate := entities.ExchangeRate{}
	//if err := json.Unmarshal(body, &exchangeRate); err != nil {
	//	log.Print(err)
	//	log.Print("failed to nmarshal exchangeRateData")
	//}

	resp, err := c.RestyClientAddress.R().
		SetQueryParams(map[string]string{
			"source":     "USD",
			"currencies": "UAH",
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("apikey", config.GetExchangeTokenAPI()).
		Get("https://api.apilayer.com/currency_data/live")

	if err != nil {
		log.Print(err)
		return exchangeRate, err
	}
	//
	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
		log.Print(err)
		log.Print("failed to nmarshal exchangeRateData")
		return exchangeRate, err
	}
	//
	//if exchangeRate.Success != true {
	//	log.Print(err)
	//	log.Print("cant get exchange rate data")
	//	return exchangeRate, err
	//}
	return exchangeRate, nil

}

//// Get exchange rate USD to UAH
//func (c ExchangeAdapter) GetExchangeRateA() (entities.ExchangeRate, error) {
//
//	resp, err := c.RestyClientAddress.R().
//		EnableTrace().
//		SetQueryParams(map[string]string{
//			"source":     "USD",
//			"currencies": "UAH",
//		}).
//		SetHeader("Accepts", "application/json").
//		SetHeader("apikey", config.GetExchangeTokenAPI()).
//		Get("https://api.currencylayer.com/live")
//
//	var exchangeRate = entities.ExchangeRate{}
//	if err != nil {
//		log.Print(err)
//		return exchangeRate, err
//	}
//
//	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
//		log.Print(err)
//		log.Print("failed to nmarshal exchangeRateData")
//		return exchangeRate, err
//	}
//
//	if exchangeRate.Success != true {
//		log.Print(err)
//		log.Print("cant get exchange rate data")
//		return exchangeRate, err
//	}
//	return exchangeRate, nil
//}
