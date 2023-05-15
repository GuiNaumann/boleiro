package players

import (
	"boleiro/domain/entities"
	"boleiro/infrastructure/repositories/players"
	"boleiro/view/http_error"
	"context"
	"database/sql"
	"log"
	"strings"
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

	player.Name = strings.TrimSpace(player.Name)

	if player.Name == "" {
		return http_error.NewBadRequestError("Nome não definido.")
	}
	if len(player.Name) > 100 {
		return http_error.NewBadRequestError("Nome não pode conter mais de 100 caracteres.")
	}

	return u.playersRepo.Create(ctx, player)
}
func (u useCases) Update(ctx context.Context, player entities.Players, playerId int64) error {
	player.Name = strings.TrimSpace(player.Name)

	if player.Name == "" {
		return http_error.NewBadRequestError("Nome não definido.")
	}
	if len(player.Name) > 100 {
		return http_error.NewBadRequestError("Nome não pode conter mais de 100 caracteres.")
	}

	_, err := u.playersRepo.GetById(ctx, playerId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[] Error ")
		return http_error.NewInternalServerError(http_error.UnexpectedError)
	}
	if err == sql.ErrNoRows {
		return http_error.NewBadRequestError("Jogador não encontrado.")
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
		log.Println("[GetById] Error GetById testeee2222 ", err)

	}
	return players, nil
}
