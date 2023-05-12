package relation

import (
	"boleiro/domain/entities"
	"context"
)

type UseCases interface {
	// Create new relation.
	Create(ctx context.Context, idPlayer int64, idSponsor int64) error
	// Delete remove a relation.
	Delete(ctx context.Context, idPlayer int64, idSponsor int64) error
	// GetById remove a relation.
	GetById(ctx context.Context, idPlayer int64) ([]entities.Sponsor, error)
	GetByIdS(ctx context.Context, idSponsor int64) ([]entities.Players, error)
}
