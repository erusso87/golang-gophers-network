package gophers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type GopherRepository interface {
	InsertGopher(gopher Gopher) error
	ListGophers() ([]Gopher, error)
	Close()
}

type dbHandler struct {
	db *sql.DB
}

func (handler *dbHandler) InsertGopher(gopher Gopher) error {
	_, err := handler.db.Exec(
		"INSERT INTO gophers(id, name, color, weight, created_at) VALUES($1, $2, $3, $4, $5)",
		gopher.ID,
		gopher.Name,
		gopher.Color,
		gopher.Weight,
		gopher.CreatedAt,
	)

	return err
}

func (handler *dbHandler) ListGophers() ([]Gopher, error) {
	rows, err := handler.db.Query("SELECT * FROM gophers ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var gophers []Gopher

	for rows.Next() {
		gopher := Gopher{}
		if err = rows.Scan(&gopher.ID, &gopher.Name, &gopher.Color, &gopher.Weight, &gopher.CreatedAt, ); err == nil {
			gophers = append(gophers, gopher)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gophers, nil
}

func (handler *dbHandler) Close() {
	handler.db.Close()
}

func CreateRepository(url string) (GopherRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &dbHandler{db}, nil
}