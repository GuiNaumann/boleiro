package players

import (
	"boleiro/domain/entities"
	"context"
)

type UseCases interface {
	// Create new players.
	Create(ctx context.Context, players entities.Players) error

	// Update  a players.
	Update(ctx context.Context, players entities.Players, userId int64) error

	// Delete remove a players.
	Delete(ctx context.Context, userId int64) error

	// GetAll return all players.
	GetAll(ctx context.Context) ([]entities.Players, error)

	// GetById return a players by id.
	GetById(ctx context.Context, userId int64) (entities.Players, error)
}
