package db

import (
	"database/sql"
	"context"
	"gophers-network/schema"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertGopher(ctx context.Context, gopher schema.Gopher) error {
	_, err := r.db.Exec(
		"INSERT INTO gophers(id, body, created_at) VALUES($1, $2, $3)",
		gopher.ID,
		gopher.Body,
		gopher.CreatedAt,
	)

	return err
}

func (r *PostgresRepository) ListGophers(ctx context.Context, offset uint64, limit uint64) ([]schema.Gopher, error) {
	rows, err := r.db.Query(
		"SELECT * FROM gophers ORDER BY id DESC OFFSET $1 LIMIT $2",
		offset,
		limit,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var gophers []schema.Gopher

	for rows.Next() {
		gopher := schema.Gopher{}
		if err = rows.Scan(&gopher.ID, &gopher.Body, &gopher.CreatedAt); err == nil {
			gophers = append(gophers, gopher)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gophers, nil
}