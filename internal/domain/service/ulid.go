package service

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULIDService struct {
}

func NewULIDService() ULIDService {
	return ULIDService{}
}

func (u *ULIDService) GenerateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
