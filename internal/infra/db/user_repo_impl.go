package db

import (
	"database/sql"
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
	jptime "inbody-ocr-backend/pkg/jp_time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user entity.User) (*entity.User, error) {
	query := `INSERT INTO users (id, name, email, password, organization_id, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	now := jptime.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.OrganizationID, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreateUserWithTx(tx bun.Tx, user entity.User) (*entity.User, error) {
	query := `INSERT INTO users (id, name, email, password, organization_id, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	now := jptime.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := tx.Exec(query, user.ID, user.Name, user.Email, user.Password, user.OrganizationID, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) HashPassword(password string) (string, error) {
	if password == "" || len(password) == 0 {
		return "", fmt.Errorf("password is required")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (r *userRepository) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	query := `SELECT id, name, email, password, organization_id, role, created_at, updated_at FROM users WHERE email = ?`

	var user entity.User
	var createdAt, updatedAt string
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.OrganizationID, &user.Role, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// ユーザーが存在しない場合はnilとエラーを返す
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	user.CreatedAt, err = jptime.ParseDateTime(createdAt)
	if err != nil {
		return nil, err
	}

	user.UpdatedAt, err = jptime.ParseDateTime(updatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(id string) (*entity.User, error) {
	query := `SELECT id, name, email, password, organization_id, role, created_at, updated_at FROM users WHERE id = ?`

	var user entity.User
	var createdAt, updatedAt string
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.OrganizationID, &user.Role, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	user.CreatedAt, err = jptime.ParseDateTime(createdAt)
	if err != nil {
		return nil, err
	}

	user.UpdatedAt, err = jptime.ParseDateTime(updatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateRole(id string, role entity.OrganizationRole) error {
	query := `UPDATE users SET role = ? WHERE id = ?`

	_, err := r.db.Exec(query, role, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UserExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`

	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) BeginTransaction() (bun.Tx, error) {
	return r.db.Begin()
}
