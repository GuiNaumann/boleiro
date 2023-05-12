package relation

import (
	"boleiro/domain/entities"
	"boleiro/infrastructure/repositories/relation"
	"context"
)

type useCases struct {
	relationRepo relation.Repository
}

func NewUseCases(relationRepo relation.Repository) UseCases {
	return &useCases{
		relationRepo: relationRepo,
	}
}

func (u useCases) Create(ctx context.Context, idPlayer int64, idSponsor int64) error {
	return u.relationRepo.Create(ctx, idPlayer, idSponsor)
}
func (u useCases) Delete(ctx context.Context, idPlayer int64, idSponsor int64) error {
	return u.relationRepo.Delete(ctx, idPlayer, idSponsor)

}
func (u useCases) GetById(ctx context.Context, idPlayer int64) ([]entities.Sponsor, error) {
	return u.relationRepo.GetById(ctx, idPlayer)
}
func (u useCases) GetByIdS(ctx context.Context, idSponsor int64) ([]entities.Players, error) {
	return u.relationRepo.GetByIdS(ctx, idSponsor)
}
