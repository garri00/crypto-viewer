package dtos

import (
	"time"
)

type Coin struct {
	ID             string    `bson:"_id,omitempty"`
	CoinID         int       `bson:"coin-id"`
	Name           string    `bson:"name"`
	Symbol         string    `bson:"symbol"`
	NumMarketPairs int       `bson:"numMarketPairs"`
	DateAdded      time.Time `bson:"dateAdded"`
	MaxSupply      float64   `bson:"maxSupply"`
	Price          float64   `bson:"price"`
	MarketCap      float64   `bson:"marketCap"`
	LastUpdated    time.Time `bson:"lastUpdated"`
}
