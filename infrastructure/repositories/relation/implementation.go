package relation

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
func (r repository) Create(ctx context.Context, idPlayer int64, idSponsor int64) error {
	query := `
	INSERT INTO relation (id_players, id_sponsor) VALUES (?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, idPlayer, idSponsor)
	if err != nil {
		log.Println("[Create] Error ExecContext", err)
		return err
	}
	return nil
}
func (r repository) Delete(ctx context.Context, idPlayer int64, idSponsor int64) error {
	//language=sql
	query := `
	UPDATE relation 
	SET status_code = ?
	WHERE id_players = ? and id_sponsor = ?
	`

	_, err := r.db.ExecContext(ctx, query, entities.StatusDeleted, idPlayer, idSponsor)
	if err != nil {
		log.Println("[Delete] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetById(ctx context.Context, idPlayer int64) ([]entities.Sponsor, error) {
	//language=sql
	query := `
	SELECT s.id,
	       s.name,
	       s.status_code,
	       s.created_at,
	       s.modified_at
	FROM sponsor s
	INNER JOIN relation ON s.id = relation.id_sponsor AND relation.id_players = ?
	WHERE s.status_code != 2`

	rows, err := r.db.QueryContext(ctx, query, idPlayer)
	if err != nil {
		log.Println("[GetById] Error QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var sponsors []entities.Sponsor
	for rows.Next() {
		var sponsor entities.Sponsor
		err = rows.Scan(
			&sponsor.Id,
			&sponsor.Name,
			&sponsor.StatusCode,
			&sponsor.CreatedAt,
			&sponsor.ModifiedAt,
		)
		if err != nil {
			log.Println("[GetById] Error Scan", err)
			return nil, err
		}
		sponsors = append(sponsors, sponsor)
	}
	return sponsors, nil
}
func (r repository) GetByIdS(ctx context.Context, idSponsor int64) ([]entities.Players, error) {
	//language=sql
	query := `
	SELECT p.id,
	       p.name,
	       p.status_code,
	       p.created_at,
	       p.modified_at
	FROM players p
	INNER JOIN relation ON p.id = relation.id_players AND relation.id_sponsor = ?
	WHERE p.status_code != 2`

	rows, err := r.db.QueryContext(ctx, query, idSponsor)
	if err != nil {
		log.Println("[GetByIdS] Error QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var player []entities.Players
	for rows.Next() {
		var players entities.Players
		err = rows.Scan(
			&players.Id, &players.Name, &players.StatusCode, &players.CreatedAt, &players.ModifiedAt,
		)
		if err != nil {
			log.Println("[GetByIdS] Error Scan", err)
			return nil, err
		}
		player = append(player, players)
	}
	return player, nil
}
