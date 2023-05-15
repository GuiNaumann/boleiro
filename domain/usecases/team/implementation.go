package team

import (
	"boleiro/domain/entities"
	"boleiro/infrastructure/repositories/team"
	"context"
	"errors"
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
		return errors.New("Nome não definido.")
	}
	if len(team.Name) > 20 {
		return errors.New("Nome não pode conter mais de 20 caracteres.")
	}

	return u.teamRepo.Create(ctx, team)
}
func (u useCases) Update(ctx context.Context, team entities.Team, teamId int64) error {
	team.Name = strings.TrimSpace(team.Name)

	if team.Name == "" {
		return errors.New("Nome não definido.")
	}
	if len(team.Name) > 20 {
		return errors.New("Nome não pode conter mais de 20 caracteres.")
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
