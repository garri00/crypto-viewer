package entities

import (
	"time"
)

type CoinsData struct {
	Coins []Coin `json:"data" bson:"coins"`
}

type USD struct {
	Price                 float64   `json:"price" bson:"price"`
	Volume24H             float64   `json:"volume_24h" bson:"volume24H"`
	VolumeChange24H       float64   `json:"volume_change_24h" bson:"volumeChange24H"`
	PercentChange1H       float64   `json:"percent_change_1h" bson:"percentChange1H"`
	PercentChange24H      float64   `json:"percent_change_24h" bson:"percentChange24H"`
	PercentChange7D       float64   `json:"percent_change_7d" bson:"percentChange7D"`
	PercentChange30D      float64   `json:"percent_change_30d" bson:"percentChange30D"`
	PercentChange60D      float64   `json:"percent_change_60d" bson:"percentChange60D"`
	PercentChange90D      float64   `json:"percent_change_90d" bson:"percentChange90D"`
	MarketCap             float64   `json:"market_cap" bson:"marketCap"`
	MarketCapDominance    float64   `json:"market_cap_dominance" bson:"marketCapDominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap" bson:"fullyDilutedMarketCap"`
	LastUpdated           time.Time `json:"last_updated" bson:"lastUpdated"`
}

type Coin struct {
	ID                int       `json:"id" bson:"_id"`
	Name              string    `json:"name" bson:"name"`
	Symbol            string    `json:"symbol" bson:"symbol"`
	Slug              string    `json:"slug" bson:"slug"`
	NumMarketPairs    int       `json:"num_market_pairs" bson:"numMarketPairs"`
	DateAdded         time.Time `json:"date_added" bson:"dateAdded"`
	MaxSupply         *int      `json:"max_supply" bson:"maxSupply"`
	CirculatingSupply float64   `json:"circulating_supply" bson:"circulatingSupply"`
	TotalSupply       float64   `json:"total_supply" bson:"totalSupply"`
	Quote             `json:"quote" bson:"quote"`
}

type Quote struct {
	USD `json:"USD" bson:"USD"`
}
