package db

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
	"time"
)

type organizationRepository struct {
	db database.DB
}

func NewOrganizationRepository(db *database.DB) repository.OrganizationRepository {
	return &organizationRepository{
		db: *db,
	}
}

// 名前被りはOK, 正し、同じユーザが同じ名前の組織を作成することはできない
func (r *organizationRepository) CreateOrganization(org entity.Organization) (*entity.Organization, error) {
	query := `INSERT INTO organizations (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)`

	now := time.Now()
	org.CreatedAt = now
	org.UpdatedAt = now

	_, err := r.db.Exec(query, org.ID, org.Name, org.CreatedAt, org.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (r *organizationRepository) FindByID(id string) (*entity.Organization, error) {
	query := `SELECT * FROM organizations WHERE id = ?`

	row := r.db.QueryRow(query, id)

	var org entity.Organization
	if err := row.Scan(&org.ID, &org.Name, &org.CreatedAt, &org.UpdatedAt); err != nil {
		return nil, err
	}

	return &org, nil
}

func (r *organizationRepository) DeleteOrganization(id string) error {
	query := `DELETE FROM organizations WHERE id = ?`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *organizationRepository) GetMember(orgID string) ([]entity.User, error) {
	query := `SELECT id, name, email, password, role, organization_id, created_at, updated_at FROM users WHERE organization_id = ?`

	rows, err := r.db.Query(query, orgID)
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		var createdAt, updatedAt string
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.OrganizationID, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
