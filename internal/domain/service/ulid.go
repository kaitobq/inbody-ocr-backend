package service

import (
	"crypto/rand"
	jptime "inbody-ocr-backend/pkg/jp_time"

	"github.com/oklog/ulid/v2"
)

type ULIDService struct {
}

func NewULIDService() ULIDService {
	return ULIDService{}
}

func (u *ULIDService) GenerateULID() string {
	t := jptime.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
