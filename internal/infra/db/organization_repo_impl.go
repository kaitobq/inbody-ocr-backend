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
