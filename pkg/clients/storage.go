package clients

import (
	"context"

	"crypto-viewer/src/entities"
	"crypto-viewer/src/entities/dtos"
)

type Storage interface {
	Create(ctx context.Context, coins dtos.Coin) error
	FindOne(ctx context.Context, id string) (c dtos.Coin, err error)
	FindAll(ctx context.Context) (c []dtos.Coin, err error)
	Update(ctx context.Context, coins entities.Coin) error
	Delete(ctx context.Context, id string) error
}
