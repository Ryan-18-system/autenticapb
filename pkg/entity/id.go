package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID = ulid.ULID

func NewID() ID {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.MustNew(ms, entropy)
}
