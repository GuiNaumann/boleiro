package players

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
func (r repository) Create(ctx context.Context, players entities.Players) error {
	query := `
	INSERT INTO players (name) VALUES (?)
	`

	_, err := r.db.ExecContext(ctx, query, players.Name)
	if err != nil {
		log.Println("[Create] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) Update(ctx context.Context, players entities.Players, playerId int64) error {
	query := `
	UPDATE players SET name = ? 
	WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, players.Name, playerId)
	if err != nil {
		log.Println("[Upadate] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetAll(ctx context.Context) ([]entities.Players, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM players`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}

	var players []entities.Players
	for rows.Next() {
		var player entities.Players
		err = rows.Scan(&player.Id, &player.Name, &player.StatusCode, &player.CreatedAt, &player.ModifiedAt)
		if err != nil {
			log.Println("[Create] Error Scan", err)
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}
func (r repository) Delete(ctx context.Context, playerId int64) error {
	//language=sql
	query := `
	UPDATE players 
	SET status_code = ?
	WHERE id = ? 
	`

	_, err := r.db.ExecContext(ctx, query, entities.StatusDeleted, playerId)
	if err != nil {
		log.Println("[Delete] Error ExecContext", err)
		return err
	}

	return nil
}
