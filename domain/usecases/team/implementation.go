package team

import (
	"boleiro/domain/entities"
	"boleiro/infrastructure/repositories/team"
	"boleiro/view/http_error"
	"context"
	"database/sql"
	"log"
	"strings"
)

type useCases struct {
	teamRepo team.Repository
}

func NewUseCases(teamRepo team.Repository) UseCases {
	return &useCases{
		teamRepo: teamRepo,
	}
}

func (u useCases) Create(ctx context.Context, team entities.Team) error {
	team.Name = strings.TrimSpace(team.Name)

	if team.Name == "" {
		return http_error.NewBadRequestError("Nome não definido.")
	}
	if len(team.Name) > 100 {
		return http_error.NewBadRequestError("Nome não pode conter mais de 100 caracteres.")
	}

	return u.teamRepo.Create(ctx, team)
}
func (u useCases) Update(ctx context.Context, team entities.Team, teamId int64) error {
	team.Name = strings.TrimSpace(team.Name)

	if team.Name == "" {
		return http_error.NewBadRequestError("Nome não definido.")
	}
	if len(team.Name) > 100 {
		return http_error.NewBadRequestError("Nome não pode conter mais de 100 caracteres.")
	}

	_, err := u.teamRepo.GetById(ctx, teamId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[Update] Error GetById", err)
		return http_error.NewInternalServerError(http_error.UnexpectedError)
	}
	if err == sql.ErrNoRows {
		return http_error.NewBadRequestError("Time não encontrado.")
	}

	return u.teamRepo.Update(ctx, team, teamId)
}
func (u useCases) Delete(ctx context.Context, teamId int64) error {
	return u.teamRepo.Delete(ctx, teamId)
}
func (u useCases) GetAll(ctx context.Context) ([]entities.Team, error) {
	return u.teamRepo.GetAll(ctx)
}
func (u useCases) GetById(ctx context.Context, userId int64) (*entities.Team, error) {
	team, err := u.teamRepo.GetById(ctx, userId)
	if err != nil {
		log.Println("[GetById] Error GetById", err)
	}
	return team, nil
}
