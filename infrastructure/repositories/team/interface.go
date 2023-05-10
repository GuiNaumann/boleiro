package team

import (
	"boleiro/domain/entities"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entities.Team, error)
	Create(ctx context.Context, team entities.Team) error
	Update(ctx context.Context, team entities.Team, teamId int64) error
	Delete(ctx context.Context, teamId int64) error
}
