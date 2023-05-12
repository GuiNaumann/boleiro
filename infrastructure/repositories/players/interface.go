package players

import (
	"boleiro/domain/entities"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context, filter entities.ListFilter) ([]entities.Players, error)
	Create(ctx context.Context, players entities.Players) error
	Update(ctx context.Context, players entities.Players, playerId int64) error
	Delete(ctx context.Context, playerId int64) error
	GetById(ctx context.Context, userId int64) (*entities.Players, error)
	GetByName(ctx context.Context, userName int64) (*entities.Players, error)
}
