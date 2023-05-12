package sponsor

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
func (r repository) Create(ctx context.Context, sponsor entities.Sponsor) error {
	query := `
	INSERT INTO sponsor (name) VALUES (?)
	`

	_, err := r.db.ExecContext(ctx, query, sponsor.Name)
	if err != nil {
		log.Println("[Create] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) Update(ctx context.Context, sponsor entities.Sponsor, sponsorId int64) error {
	query := `
	UPDATE sponsor SET name = ? 
	WHERE id = ? AND 
	      status_code != ?
	`

	_, err := r.db.ExecContext(ctx, query, sponsor.Name, sponsorId, entities.StatusDeleted)
	if err != nil {
		log.Println("[Upadate] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetAll(ctx context.Context) ([]entities.Sponsor, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM sponsor
	WHERE status_code != ?`

	rows, err := r.db.QueryContext(ctx, query, entities.StatusDeleted)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}

	var sponsor []entities.Sponsor
	for rows.Next() {
		var sponsors entities.Sponsor
		err = rows.Scan(&sponsors.Id, &sponsors.Name, &sponsors.StatusCode, &sponsors.CreatedAt, &sponsors.ModifiedAt)
		if err != nil {
			log.Println("[Create] Error Scan", err)
			return nil, err
		}
		sponsor = append(sponsor, sponsors)
	}
	defer rows.Close()

	return sponsor, nil
}
func (r repository) Delete(ctx context.Context, sponsorId int64) error {
	//language=sql
	query := `
	UPDATE sponsor 
	SET status_code = ?
	WHERE id = ? 
	`

	_, err := r.db.ExecContext(ctx, query, entities.StatusDeleted, sponsorId)
	if err != nil {
		log.Println("[Delete] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetById(ctx context.Context, idSponsor int64) (*entities.Sponsor, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM sponsor
	WHERE id = ? AND 
	      status_code = 0`

	var sponsors entities.Sponsor
	err := r.db.QueryRowContext(ctx, query, idSponsor).Scan(&sponsors.Id, &sponsors.Name, &sponsors.StatusCode, &sponsors.CreatedAt, &sponsors.ModifiedAt)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}
	return &sponsors, err
}
