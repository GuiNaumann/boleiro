package sponsor

import (
	"boleiro/domain/entities"
	"context"
)

type UseCases interface {
	// Create new sponsor.
	Create(ctx context.Context, sponsor entities.Sponsor) error

	// Update  a sponsor.
	Update(ctx context.Context, sponsor entities.Sponsor, userId int64) error

	// Delete remove a sponsor.
	Delete(ctx context.Context, userId int64) error

	// GetAll return all sponsor.
	GetAll(ctx context.Context) ([]entities.Sponsor, error)

	// GetById return a sponsor by id.
	GetById(ctx context.Context, userId int64) (entities.Sponsor, error)
}
