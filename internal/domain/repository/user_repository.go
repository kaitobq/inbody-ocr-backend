package repository

import "inbody-ocr-backend/internal/domain/entity"

type UserRepository interface {
	CheckDuplicateEmail(email string) (bool, error)
	CreateUser(user entity.User) error
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}
