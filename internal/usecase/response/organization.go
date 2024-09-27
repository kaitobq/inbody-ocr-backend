package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"time"
)

type OrganizationResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrganizationResponse struct {
	Message      string               `json:"message"`
	Organization OrganizationResponse `json:"organization"`
	User         UserResponse         `json:"user"`
	Token        TokenResponse        `json:"token"`
}

func NewCreateOrganizationResponse(org entity.Organization, token, userID, userName string, exp time.Time) (*CreateOrganizationResponse, error) {
	return &CreateOrganizationResponse{
		Message: "Organization created successfully",
		Organization: OrganizationResponse{
			ID:        org.ID,
			Name:      org.Name,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		},
		User: UserResponse{
			ID:   userID,
			Name: userName,
			Role: entity.OrganizationRoleOwner,
		},
		Token: TokenResponse{
			Value: token,
			Exp:   exp,
		},
	}, nil
}
