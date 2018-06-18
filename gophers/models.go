package gophers

import (
	"time"
)

type Gopher struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
	Weight float32 `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}