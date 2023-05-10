package players

import (
	"boleiro/domain/entities"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entities.Players, error)
	Create(ctx context.Context, players entities.Players) error
	Update(ctx context.Context, players entities.Players, playerId int64) error
	Delete(ctx context.Context, playerId int64) error
}
