package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"net/http"
	"time"
)

type OrganizationResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrganizationResponse struct {
	Status       int                  `json:"status"`
	Message      string               `json:"message"`
	Organization OrganizationResponse `json:"organization"`
	User         UserResponse         `json:"user"`
	Token        TokenResponse        `json:"token"`
}

func NewCreateOrganizationResponse(org entity.Organization, token, userID, userName string, exp time.Time) (*CreateOrganizationResponse, error) {
	return &CreateOrganizationResponse{
		Status:  http.StatusCreated,
		Message: "ok",
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
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Members []UserResponse `json:"members"`
}

func NewGetAllMembersResponse(users []entity.User) (*GetAllMembersResponse, error) {
	var members []UserResponse
	for _, user := range users {
		members = append(members, *NewUserResponse(user))
	}

	return &GetAllMembersResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Members: members,
	}, nil
}

type UpdateRoleResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Updated UserResponse `json:"updated"`
}

func NewUpdateRoleResponse(user entity.User) (*UpdateRoleResponse, error) {
	return &UpdateRoleResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Updated: *NewUserResponse(user),
	}, nil
}

type DeleteMemberResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewDeleteMemberResponse() (*DeleteMemberResponse, error) {
	return &DeleteMemberResponse{
		Status:  http.StatusOK,
		Message: "ok",
	}, nil
}
