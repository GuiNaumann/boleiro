package relation

import (
	"boleiro/domain/entities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, idPlayer int64, idSponsor int64) error
	Delete(ctx context.Context, idPlayer int64, idSponsor int64) error
	GetById(ctx context.Context, idPlayer int64) ([]entities.Sponsor, error)
	GetByIdS(ctx context.Context, idSponsor int64) ([]entities.Players, error)
}
