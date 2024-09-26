package repository

import "inbody-ocr-backend/internal/domain/entity"

type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
	DeleteUser(id string) error
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	UserExists(email string) (bool, error)
}
