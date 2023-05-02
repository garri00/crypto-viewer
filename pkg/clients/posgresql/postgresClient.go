package posgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

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
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", configs.Username, configs.Password, configs.Host, configs.Port, configs.Database)

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

	//var embedMigrations embed.FS
	//goose.SetBaseFS(embedMigrations)
	//
	////gooseDB, err := sql.Open("pgx", connectionString)
	////if err != nil {
	////	logger.Log.Err(err).Msg("failed to connect postgreDB")
	////	return nil, err
	////}
	//
	//db2, err := sql.Open("pgx", "user=msylniahin password=zxc1212 host=localhost port=5432 database=crypto sslmode=disable")
	//if err != nil {
	//	return nil, err
	//}
	//
	//if err := goose.SetDialect("postgres"); err != nil {
	//	logger.Log.Err(err).Msg("goose can't select dialog")
	//	return nil, err
	//}
	//
	//if err := goose.Up(db2, "migrations"); err != nil {
	//	logger.Log.Err(err).Msg("migrations up failed")
	//	return nil, err
	//}

	return dbpool, nil
}
