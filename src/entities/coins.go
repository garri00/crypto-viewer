package entities

import (
	"time"
)

type CoinsData struct {
	Coins []struct {
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
