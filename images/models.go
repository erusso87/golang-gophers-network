package images

import "time"

type Image struct {
	ID string `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}