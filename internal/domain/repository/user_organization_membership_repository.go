package repository

import "inbody-ocr-backend/internal/domain/entity"

type UserOrganizationMembershipRepository interface {
	CreateMembership(entity.UserOrganizationMembership) (*entity.UserOrganizationMembership, error)
	FindByUserID(userID string) (*[]entity.UserOrganizationMembership, error)
	DeleteMembership(userID, orgID string) error
	GetRole(userID, orgID string) (string, error)
}
