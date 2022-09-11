package handlers

import (
	"crypto-viewer/src/config"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CoinsResty(w http.ResponseWriter, r *http.Request) {

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"start": r.URL.Query().Get("start"),
			"limit": r.URL.Query().Get("limit"),
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", config.TokenAPI).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET request"))
		return
	}

	fmt.Println(resp.Request.URL)

	if resp.StatusCode() != http.StatusOK {
		var errResponse struct {
			Status struct {
				ErrorCode    int    `json:"error_code"`
				ErrorMessage string `json:"error_message"`
			} `json:"status"`
		}

		if err := json.Unmarshal(resp.Body(), &errResponse); err != nil {
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

	type Coin struct {
		Data []struct {
			Id                int       `json:"id"`
			Name              string    `json:"name"`
			Symbol            string    `json:"symbol"`
			Slug              string    `json:"slug"`
			NumMarketPairs    int       `json:"num_market_pairs"`
			DateAdded         time.Time `json:"date_added"`
			MaxSupply         *int      `json:"max_supply"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			Quote             struct {
				USD struct {
					Price                 float64   `json:"price"`
					Volume24H             float64   `json:"volume_24h"`
					VolumeChange24H       float64   `json:"volume_change_24h"`
					PercentChange1H       float64   `json:"percent_change_1h"`
					PercentChange24H      float64   `json:"percent_change_24h"`
					PercentChange7D       float64   `json:"percent_change_7d"`
					PercentChange30D      float64   `json:"percent_change_30d"`
					PercentChange60D      float64   `json:"percent_change_60d"`
					PercentChange90D      float64   `json:"percent_change_90d"`
					MarketCap             float64   `json:"market_cap"`
					MarketCapDominance    float64   `json:"market_cap_dominance"`
					FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
					LastUpdated           time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"data"`
	}

	okResponse := Coin{}

	if err := json.Unmarshal(resp.Body(), &okResponse); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to unmarshal okResponse"))
		return
	}

	file, err := json.MarshalIndent(okResponse, "", " ")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to marshal okResponse"))
		return
	}
	ioutil.WriteFile("src/pkg/coinslist.json", file, 0644)
	fmt.Println(resp.Status)
	w.WriteHeader(http.StatusOK)

	w.Write(file)
}
