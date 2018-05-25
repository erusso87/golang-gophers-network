package db

import (
	"context"
	"gophers-network/schema"
)

type Repository interface {
	Close()
	InsertGopher(ctx context.Context, gopher schema.Gopher) error
	ListGophers(ctx context.Context, offset uint64, limit uint64) ([]schema.Gopher, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertGopher(ctx context.Context, gopher schema.Gopher) error {
	return impl.InsertGopher(ctx, gopher)
}

func ListGophers(ctx context.Context, offset uint64, limit uint64) ([]schema.Gopher, error) {
	return impl.ListGophers(ctx, offset, limit)
}

