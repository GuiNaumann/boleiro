package sponsor

import (
	"boleiro/domain/entities"
	"boleiro/infrastructure/repositories/sponsor"
	"context"
	"errors"
	"log"
	"strings"
)

type useCases struct {
	sponsorRepo sponsor.Repository
}

func NewUseCases(sponsorRepo sponsor.Repository) UseCases {
	return &useCases{
		sponsorRepo: sponsorRepo,
	}
}
func (u useCases) Create(ctx context.Context, sponsor entities.Sponsor) error {
	sponsor.Name = strings.TrimSpace(sponsor.Name)

	if sponsor.Name == "" {
		return errors.New("Nome não definido.")
	}
	if len(sponsor.Name) > 20 {
		return errors.New("Nome não pode conter mais de 20 caracteres.")
	}

	return u.sponsorRepo.Create(ctx, sponsor)

}
func (u useCases) Update(ctx context.Context, sponsor entities.Sponsor, sponsorId int64) error {
	sponsor.Name = strings.TrimSpace(sponsor.Name)

	if sponsor.Name == "" {
		return errors.New("Nome não definido.")
	}
	if len(sponsor.Name) > 20 {
		return errors.New("Nome não pode conter mais de 20 caracteres.")
	}

	return u.sponsorRepo.Update(ctx, sponsor, sponsorId)

}
func (u useCases) Delete(ctx context.Context, sponsorId int64) error {
	return u.sponsorRepo.Delete(ctx, sponsorId)
}
func (u useCases) GetAll(ctx context.Context, filter entities.ListFilter) ([]entities.Sponsor, error) {
	return u.sponsorRepo.GetAll(ctx, filter)
}
func (u useCases) GetById(ctx context.Context, userId int64) (*entities.Sponsor, error) {

	sponsor, err := u.sponsorRepo.GetById(ctx, userId)
	if err != nil {
		log.Println("[GetById] Error GetById", err)
	}
	return sponsor, nil
}
