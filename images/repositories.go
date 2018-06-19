package images

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type ImageRepository interface {
	InsertImage(image Image) error
	GetImage(id string) *Image
	Close()
}

type dbHandler struct {
	db *sql.DB
}

func (handler *dbHandler) InsertImage(image Image) error {
	_, err := handler.db.Exec(
		"INSERT INTO images(id, content, created_at) VALUES($1, $2, $3)",
		image.ID,
		image.Content,
		image.CreatedAt,
	)

	return err
}

func (handler *dbHandler) GetImage(id string) *Image {
	row := handler.db.QueryRow("SELECT * FROM images WHERE id = $1", id)
	if row == nil {
		return nil
	}

	var image = Image{}
	if err := row.Scan(&image.ID, &image.Content, &image.CreatedAt); err != nil {
		return nil
	}

	return &image
}

func (handler *dbHandler) Close() {
	handler.db.Close()
}

func CreateRepository(url string) (ImageRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &dbHandler{db}, nil
}