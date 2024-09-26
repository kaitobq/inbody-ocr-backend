package entity

import "time"

type OrganizationRole string

const (
	OrganizationRoleOwner  OrganizationRole = "owner"
	OrganizationRoleAdmin  OrganizationRole = "admin"
	OrganizationRoleMember OrganizationRole = "member"
)

type User struct {
	ID             string
	Name           string
	Email          string
	Password       string
	OrganizationID string
	Role           OrganizationRole
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
