package usecases

import (
	"context"

	"github.com/rs/zerolog"

	"crypto-viewer/pkg/clients"
	"crypto-viewer/src/entities"
	"crypto-viewer/src/entities/dtos"
)

type SaveCoinsDB struct {
	log zerolog.Logger
	db  clients.Storage
}

func NewSaveCoinsDB(save clients.Storage, l zerolog.Logger) SaveCoinsDB {
	return SaveCoinsDB{
		db:  save,
		log: l,
	}
}

func (s SaveCoinsDB) SaveCoinsDB(coinsData entities.CoinsData) error {
	for i := 0; i < len(coinsData.Coins); i++ {
		coin := dtos.Coin{
			CoinID:         coinsData.Coins[i].ID,
			Name:           coinsData.Coins[i].Name,
			Symbol:         coinsData.Coins[i].Symbol,
			NumMarketPairs: coinsData.Coins[i].NumMarketPairs,
			DateAdded:      coinsData.Coins[i].DateAdded,
			MaxSupply:      coinsData.Coins[i].MaxSupply,
			Price:          coinsData.Coins[i].Price,
			MarketCap:      coinsData.Coins[i].MarketCap,
			LastUpdated:    coinsData.Coins[i].LastUpdated,
		}

		err := s.db.Create(context.Background(), coin)
		if err != nil {
			s.log.Err(err).Msg("failed to save coins data")

			return err
		}
	}

	return nil
}
