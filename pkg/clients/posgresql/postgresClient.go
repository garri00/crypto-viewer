package posgresql

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/config"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, configs config.PostgreDBConf) (db *pgxpool.Pool, err error) {
	var connectionString string
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", configs.Username, configs.Password, configs.Host, configs.Port, configs.Database)

	dbpool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect postgreDB")
		return nil, err
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		logger.Log.Err(err).Msg("failed to ping postgreDB")
		return nil, err
	}

	log.Info().Msg("successfully connected to PostgresDB")

	m, err := migrate.New("file://src/migrations", connectionString)
	if err != nil {
		logger.Log.Err(err).Msg("failed to find migrations")
		return nil, err
	}

	//TODO: Як правильно тут обробити помилку бо при вже наявній міграції видає помилку no changes
	if err := m.Up(); err != nil {
		logger.Log.Err(err).Msg("migration up with err")
	}

	version, _, err := m.Version()
	log.Info().Msgf("database migrated to ver %v", version)

	return dbpool, nil
}
