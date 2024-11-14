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
	Organization OrganizationResponse `json:"organization"`
	User         UserResponse         `json:"user"`
	Token        TokenResponse        `json:"token"`
}

func NewCreateOrganizationResponse(org entity.Organization, token, userID, userName string, exp time.Time) (*CreateOrganizationResponse, error) {
	return &CreateOrganizationResponse{
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

type GetAllMembersResponse struct {
	Members []UserResponse `json:"members"`
}

func NewGetAllMembersResponse(users []entity.User) (*GetAllMembersResponse, error) {
	var members []UserResponse
	for _, user := range users {
		members = append(members, *NewUserResponse(user))
	}

	return &GetAllMembersResponse{
		Members: members,
	}, nil
}

type UpdateRoleResponse struct {
	Updated UserResponse `json:"updated"`
}

func NewUpdateRoleResponse(user entity.User) (*UpdateRoleResponse, error) {
	return &UpdateRoleResponse{
		Updated: *NewUserResponse(user),
	}, nil
}
