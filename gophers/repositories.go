package gophers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Create a generic repository structure
type GopherRepository interface {
	InsertGopher(gopher Gopher) error
	UpdateGopher(gopher Gopher) error
	ListGophers() ([]Gopher, error)
	GetGopher(id string) *Gopher
	Close()
}

type dbHandler struct {
	db *sql.DB
}

func (handler *dbHandler) InsertGopher(gopher Gopher) error {
	_, err := handler.db.Exec(
		"INSERT INTO gophers(id, name, color, weight, image, created_at) VALUES($1, $2, $3, $4, $5, $6)",
		gopher.ID,
		gopher.Name,
		gopher.Color,
		gopher.Weight,
		gopher.Image,
		gopher.CreatedAt,
	)

	return err
}

func (handler *dbHandler) UpdateGopher(gopher Gopher) error {
	_, err := handler.db.Exec(
		"UPDATE gophers SET name=$1, color=$2, weight=$3, image=$4 where id=$5",
		gopher.Name,
		gopher.Color,
		gopher.Weight,
		gopher.Image,
		gopher.ID,
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
		if err = rows.Scan(&gopher.ID, &gopher.Name, &gopher.Color, &gopher.Weight, &gopher.Image, &gopher.CreatedAt); err == nil {
			gophers = append(gophers, gopher)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gophers, nil
}

func (handler *dbHandler) GetGopher(id string) *Gopher {
	row := handler.db.QueryRow("SELECT * FROM gophers WHERE id = $1", id)
	if row == nil {
		return nil
	}

	var gopher = Gopher{}
	// TODO: Improve mass assignment
	if err := row.Scan(&gopher.ID, &gopher.Name, &gopher.Color, &gopher.Weight, &gopher.Image, &gopher.CreatedAt); err != nil {
		return nil
	}

	return &gopher
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