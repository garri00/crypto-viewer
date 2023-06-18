package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"

	"crypto-viewer/pkg/clients"
	"crypto-viewer/pkg/clients/posgresql"
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

func (s storagePG) Create(ctx context.Context, coin dtos.Coin) error {
	q := `
		INSERT INTO crypto.coins
		    (coinid,
		     coinname,
		     symbol,
		     nummarketpairs,
		     dateadded,
		     maxsupply,
		     price,
		     marketcap,
		     lastupdated)
		VALUES
		       ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	if err := s.client.QueryRow(ctx, q,
		coin.CoinID,
		coin.Name,
		coin.Symbol,
		coin.NumMarketPairs,
		coin.DateAdded,
		coin.MaxSupply,
		coin.Price,
		coin.MarketCap,
		coin.LastUpdated).Scan(&coin.ID); err != nil {
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

			s.logger.Err(newErr).Msg("error creating coin")

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
		s.logger.Err(err).Msg("error FindOne coin")

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
			s.logger.Err(err).Msg("failed to scan coin query")

			return nil, err
		}

		coins = append(coins, coin)
	}

	if err = rows.Err(); err != nil {
		s.logger.Err(err).Msg("error FindAll coins")

		return nil, err
	}

	return coins, nil
}

func (s storagePG) Update(ctx context.Context, c dtos.Coin) error {
	q := `UPDATE crypto.coins 
		   SET coinname = $1,
		     symbol = $2,
		     nummarketpairs = $3,
		     dateadded = $4,
		     maxsupply = $5,
		     price = $6,
		     marketcap = $7,
		     lastupdated = $8
           WHERE id = $9;
			`

	_, err := s.client.Query(ctx, q, c.Name, c.Symbol, c.NumMarketPairs, c.DateAdded, c.MaxSupply, c.Price, c.MarketCap, c.LastUpdated, c.ID)
	if err != nil {
		s.logger.Err(err).Msg("error to Update coin")

		return err
	}

	return nil
}

func (s storagePG) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM crypto.coins WHERE id=$1;
	`

	commandTag, err := s.client.Exec(ctx, q, id)
	if err != nil {
		s.logger.Err(err).Msg("error to Delete coin")

		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}

	s.logger.Info().Msgf("coin with id = %s sucsefuly DELETED", id)

	return nil
}
