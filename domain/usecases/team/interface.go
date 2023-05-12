package team

import (
	"boleiro/domain/entities"
	"context"
)

type UseCases interface {
	// Create new players.
	Create(ctx context.Context, team entities.Team) error

	// Update  a players.
	Update(ctx context.Context, team entities.Team, userId int64) error

	// Delete remove a players.
	Delete(ctx context.Context, userId int64) error

	// GetAll return all players.
	GetAll(ctx context.Context) ([]entities.Team, error)

	// GetById return a players by id.
	GetById(ctx context.Context, userId int64) (*entities.Team, error)
}
