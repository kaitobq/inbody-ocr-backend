package repository

import "inbody-ocr-backend/internal/domain/entity"

type OrganizationRepository interface {
	CreateOrganization(org entity.Organization) (*entity.Organization, error)
	FindByID(id string) (*entity.Organization, error)
	DeleteOrganization(id string) error
	GetMember(orgID string) ([]entity.User, error)
}
