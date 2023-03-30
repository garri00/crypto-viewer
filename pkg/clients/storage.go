package clients

import (
	"context"

	"crypto-viewer/src/entities"
)

type Storage interface {
	Create(ctx context.Context, coins entities.Coin) (string, error)
	Read(ctx context.Context, id string) (c entities.Coin, err error)
	Update(ctx context.Context, coins entities.Coin) (string, error)
	Delete(ctx context.Context, coins entities.Coin) (string, error)
}
