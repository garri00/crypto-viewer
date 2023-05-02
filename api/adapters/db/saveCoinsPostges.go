package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"

	"crypto-viewer/pkg/clients"
	"crypto-viewer/pkg/clients/posgresql"
	"crypto-viewer/src/entities"
	"crypto-viewer/src/entities/dtos"
)

type storagePG struct {
	client posgresql.Client
	logger zerolog.Logger
}

func NewStoragePG(client posgresql.Client, logger zerolog.Logger) clients.Storage {
	return &storagePG{
		client: client,
		logger: logger,
	}
}

func (r storagePG) Create(ctx context.Context, coin dtos.Coin) error {
	q := `
		INSERT INTO crypto.coins
		    (coinname,
		     symbol,
		     nummarketpairs,
		     dateadded,
		     maxsupply,
		     price,
		     marketcap,
		     lastupdated)
		VALUES
		       ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	r.logger.Info().Msg(q)

	if err := r.client.QueryRow(ctx, q, coin.Name, coin.Symbol, coin.NumMarketPairs, coin.DateAdded, coin.MaxSupply, coin.Price, coin.MarketCap, coin.LastUpdated).Scan(&coin.ID); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf(
				"SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState()))

			r.logger.Err(newErr).Msg("error creating coin")

			return newErr
		}

		return nil
	}

	return nil
}

func (s storagePG) FindOne(ctx context.Context, id string) (c dtos.Coin, err error) {
	q := `
		SELECT id, coinid, coinname, symbol, nummarketpairs, dateadded, maxsupply, price, marketcap, lastupdated FROM crypto.coins WHERE id = $1;
	`

	var coin dtos.Coin
	err = s.client.QueryRow(ctx, q, id).Scan(&coin.ID, &coin.Name, &coin.Symbol, &coin.NumMarketPairs, &coin.DateAdded, &coin.MaxSupply, &coin.Price, &coin.MarketCap, &coin.LastUpdated)
	if err != nil {
		return dtos.Coin{}, err
	}

	return coin, nil
}

func (s storagePG) FindAll(ctx context.Context) (c []dtos.Coin, err error) {
	q := `
		SELECT id, coinid, coinname, symbol, nummarketpairs, dateadded, maxsupply, price, marketcap, lastupdated FROM crypto.coins;
	`

	rows, err := s.client.Query(ctx, q)
	if err != nil {
		s.logger.Err(err).Msg("failed to execute query")
		return nil, err
	}

	coins := make([]dtos.Coin, 0)

	for rows.Next() {
		var coin dtos.Coin

		err = rows.Scan(&coin.ID, &coin.CoinID, &coin.Name, &coin.Symbol, &coin.NumMarketPairs, &coin.DateAdded, &coin.MaxSupply, &coin.Price, &coin.MarketCap, &coin.LastUpdated)
		if err != nil {
			return nil, err
		}

		coins = append(coins, coin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return coins, nil
}

func (s storagePG) Update(ctx context.Context, coins entities.Coin) error {
	//TODO implement me
	panic("implement me")
}

func (s storagePG) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
