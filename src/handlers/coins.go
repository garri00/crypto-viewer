package handlers

import (
	"crypto-viewer/src/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
}

func Coins(w http.ResponseWriter, r *http.Request) {

	//u, _ := url.ParseRequestURI("")

	u := url.URL{}
	values := u.Query()
	values.Add("start", r.URL.Query().Get("start"))
	values.Add("limit", r.URL.Query().Get("limit"))

	//values.Add("symbol", r.URL.Query().Get("symbol"))
	u.RawQuery = values.Encode()

	//u.Query().Add("start", r.URL.Query().Get("start"))
	//u.Query().Add("limit", r.URL.Query().Get("limit"))
	//u.Query().Add("convert", r.URL.Query().Get("convert"))
	//u.Query().Add("symbol", r.URL.Query().Get("symbol"))

	fmt.Println(u.String())

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"+u.String(), nil)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET request"))

		return
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", config.GetConfigTokenAPI())

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error sending request to server"))
		return
	}

	fmt.Println(resp.Request.URL)
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errResponse struct {
			Status struct {
				ErrorCode    int    `json:"error_code"`
				ErrorMessage string `json:"error_message"`
			} `json:"status"`
		}

		if err := json.Unmarshal(respBody, &errResponse); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to unmarshal errResponse"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Code=%d, Message=%s", errResponse.Status.ErrorCode, errResponse.Status.ErrorMessage)))
		log.Print(errResponse.Status)
		return
	}

	var okResponse struct {
		Coin struct {
			Status struct {
				Timestamp    time.Time   `json:"timestamp"`
				ErrorCode    int         `json:"error_code"`
				ErrorMessage interface{} `json:"error_message"`
				Elapsed      int         `json:"elapsed"`
				CreditCount  int         `json:"credit_count"`
				Notice       interface{} `json:"notice"`
				TotalCount   int         `json:"total_count"`
			} `json:"status"`
			Data struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				Slug   string `json:"slug"`

				Quote struct {
					USD struct {
						Price                 float64     `json:"price"`
						Volume24H             float64     `json:"volume_24h"`
						VolumeChange24H       float64     `json:"volume_change_24h"`
						PercentChange1H       float64     `json:"percent_change_1h"`
						PercentChange24H      float64     `json:"percent_change_24h"`
						PercentChange7D       float64     `json:"percent_change_7d"`
						PercentChange30D      float64     `json:"percent_change_30d"`
						PercentChange60D      float64     `json:"percent_change_60d"`
						PercentChange90D      float64     `json:"percent_change_90d"`
						MarketCap             float64     `json:"market_cap"`
						MarketCapDominance    float64     `json:"market_cap_dominance"`
						FullyDilutedMarketCap float64     `json:"fully_diluted_market_cap"`
						Tvl                   interface{} `json:"tvl"`
						LastUpdated           time.Time   `json:"last_updated"`
					} `json:"USD"`
				} `json:"quote"`
			} `json:"data"`
		}
	}

	if err := json.Unmarshal(respBody, &okResponse); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to unmarshal okResponse"))
		return
	}

	file, _ := json.Marshal(okResponse.Coin.Data)

	ioutil.WriteFile("scr/pkg/coinslist.json", file, 0644)
	fmt.Println(resp.Status)
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
