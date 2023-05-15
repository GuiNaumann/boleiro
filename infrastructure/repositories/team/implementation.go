package team

import (
	"boleiro/domain/entities"
	"context"
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
func (r repository) Create(ctx context.Context, team entities.Team) error {
	query := `
	INSERT INTO team (name) VALUES (?)
	`

	_, err := r.db.ExecContext(ctx, query, team.Name)
	if err != nil {
		log.Println("[Create] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) Update(ctx context.Context, team entities.Team, teamId int64) error {
	query := `
	UPDATE team SET name = ? 
	WHERE id = ? AND 
	      status_code != ?
	`

	_, err := r.db.ExecContext(ctx, query, team.Name, teamId, entities.StatusDeleted)
	if err != nil {
		log.Println("[Upadate] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetAll(ctx context.Context) ([]entities.Team, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM team 
	WHERE status_code != ?`

	rows, err := r.db.QueryContext(ctx, query, entities.StatusDeleted)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var team []entities.Team
	for rows.Next() {
		var teams entities.Team

		err = rows.Scan(&teams.Id, &teams.Name, &teams.StatusCode, &teams.CreatedAt, &teams.ModifiedAt)
		if err != nil {
			log.Println("[Create] Error Scan", err)
			return nil, err
		}
		team = append(team, teams)
	}
	return team, nil
}
func (r repository) Delete(ctx context.Context, teamId int64) error {
	//language=sql
	query := `
	UPDATE team 
	SET status_code = ?
	WHERE id = ? 
	`

	_, err := r.db.ExecContext(ctx, query, entities.StatusDeleted, teamId)
	if err != nil {
		log.Println("[Delete] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetById(ctx context.Context, idTeam int64) (*entities.Team, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM team
	WHERE id = ? AND 
	      status_code = 0  `

	var teams entities.Team
	err := r.db.QueryRowContext(ctx, query, idTeam).Scan(&teams.Id, &teams.Name, &teams.StatusCode, &teams.CreatedAt, &teams.ModifiedAt)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}
	return &teams, err
}
