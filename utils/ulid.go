package utils

import (
	"github.com/oklog/ulid"
	"time"
	"math/rand"
)

func CreateULID() (ulid.ULID) {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}