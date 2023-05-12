package players

import (
	"boleiro/domain/entities"
	"boleiro/infrastructure/repositories/players"
	"context"
	"errors"
	"log"
)

type useCases struct {
	playersRepo players.Repository
}

func NewUseCases(playersRepo players.Repository) UseCases {
	return &useCases{
		playersRepo: playersRepo,
	}
}
func (u useCases) Create(ctx context.Context, player entities.Players) error {
	if len(player.Name) > 20 {
		return errors.New("Nome não pode conter mais de 20 caracteres.")
	}

	return u.playersRepo.Create(ctx, player)
}
func (u useCases) Update(ctx context.Context, player entities.Players, playerId int64) error {
	if len(player.Name) > 20 {
		return errors.New("Nome não pode conter mais de 20 caracteres.")
	}

	return u.playersRepo.Update(ctx, player, playerId)
}
func (u useCases) Delete(ctx context.Context, playerId int64) error {
	return u.playersRepo.Delete(ctx, playerId)
}
func (u useCases) GetAll(ctx context.Context, filter entities.ListFilter) ([]entities.Players, error) {
	return u.playersRepo.GetAll(ctx, filter)
}
func (u useCases) GetById(ctx context.Context, userId int64) (*entities.Players, error) {
	players, err := u.playersRepo.GetById(ctx, userId)
	if err != nil {
		log.Println("[GetById] Error GetById", err)

	}
	return players, nil
}
