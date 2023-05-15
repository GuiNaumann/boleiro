package sponsor

import (
	"boleiro/domain/entities"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context, filter entities.ListFilter) ([]entities.Sponsor, error)
	Create(ctx context.Context, sponsor entities.Sponsor) error
	Update(ctx context.Context, sponsor entities.Sponsor, sponsorId int64) error
	Delete(ctx context.Context, sponsorId int64) error
	GetById(ctx context.Context, userId int64) (*entities.Sponsor, error)
}
