package schema

import "time"

type Gopher struct {
	ID string `json:"id"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

