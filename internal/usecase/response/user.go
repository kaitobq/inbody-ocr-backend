package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"time"
)

type UserResponse struct {
	ID        string                  `json:"id"`
	Name      string                  `json:"name"`
	Role      entity.OrganizationRole `json:"role"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func NewUserResponse(user entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type TokenResponse struct {
	Value string    `json:"value"`
	Exp   time.Time `json:"expires_at"`
}

type SignUpResponse struct {
	Token TokenResponse `json:"token"`
	User  UserResponse  `json:"user"`
}

func NewSignUpResponse(token string, exp *time.Time, user entity.User) (*SignUpResponse, error) {
	return &SignUpResponse{
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		User: *NewUserResponse(user),
	}, nil
}

type SignInResponse struct {
	Token          TokenResponse `json:"token"`
	OrganizationID string        `json:"organization_id"`
	User           UserResponse  `json:"user"`
}

func NewSignInResponse(token string, exp *time.Time, user entity.User) (*SignInResponse, error) {
	return &SignInResponse{
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		OrganizationID: user.OrganizationID,
		User:           *NewUserResponse(user),
	}, nil
}

type GetOwnInfoResponse struct {
	User UserResponse `json:"user"`
}

func NewGetOwnInfoResponse(user entity.User) (*GetOwnInfoResponse, error) {
	return &GetOwnInfoResponse{
		User: *NewUserResponse(user),
	}, nil
}
