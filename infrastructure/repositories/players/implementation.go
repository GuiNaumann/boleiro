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
	WHERE id = ? AND 
	      status_code != ?
	`

	_, err := r.db.ExecContext(ctx, query, players.Name, playerId, entities.StatusDeleted)
	if err != nil {
		log.Println("[Upadate] Error ExecContext", err)
		return err
	}

	return nil
}
func (r repository) GetAll(ctx context.Context, filter entities.ListFilter) ([]entities.Players, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM players
	WHERE status_code != ?
	`

	if filter.OrderBy != "" {
		var columnName string
		switch filter.OrderBy {
		case "id":
			columnName = "id"
		case "name":
			columnName = "name"
		case "createdAt":
			columnName = "created_at"
		case "modifiedAt":
			columnName = "modified_at"
		}

		query += ` ORDER BY ` + columnName
	}

	if filter.OrderType != "" {
		var teste string
		switch filter.OrderType {
		case "id":
			teste = "id"
		case "name":
			teste = "name"
		case "createdAt":
			teste = "created_at"
		case "modifiedAt":
			teste = "modified_at"
		}

		query += ` ORDER BY ` + teste
	}

	rows, err := r.db.QueryContext(ctx, query, entities.StatusDeleted)
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
	defer rows.Close()
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
func (r repository) GetById(ctx context.Context, playerId int64) (*entities.Players, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       status_code,
	       created_at,
	       modified_at
	FROM players
	WHERE id = ? AND 
	      status_code = 0`

	var player entities.Players
	err := r.db.QueryRowContext(ctx, query, playerId).Scan(&player.Id, &player.Name, &player.StatusCode, &player.CreatedAt, &player.ModifiedAt)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}
	return &player, err
}
