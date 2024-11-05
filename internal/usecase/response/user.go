package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"net/http"
	"time"
)

type UserResponse struct {
	ID   string                  `json:"id"`
	Name string                  `json:"name"`
	Role entity.OrganizationRole `json:"role"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func NewUserResponse(user entity.User) *UserResponse {
	return &UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type TokenResponse struct {
	Value string    `json:"value"`
	Exp   time.Time `json:"expires_at"`
}

type SignUpResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Token   TokenResponse `json:"token"`
	User    UserResponse  `json:"user"`
}

func NewSignUpResponse(token string, exp *time.Time, user entity.User) (*SignUpResponse, error) {
	return &SignUpResponse{
		Status:  http.StatusCreated,
		Message: "ok",
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		User: *NewUserResponse(user),
	}, nil
}

type SignInResponse struct {
	Status         int           `json:"status"`
	Message        string        `json:"message"`
	Token          TokenResponse `json:"token"`
	OrganizationID string        `json:"organization_id"`
	User           UserResponse  `json:"user"`
}

func NewSignInResponse(token string, exp *time.Time, user entity.User) (*SignInResponse, error) {
	return &SignInResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		OrganizationID: user.OrganizationID,
		User: *NewUserResponse(user),
	}, nil
}

type GetOwnInfoResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

func NewGetOwnInfoResponse(user entity.User) (*GetOwnInfoResponse, error) {
	return &GetOwnInfoResponse{
		Status:  http.StatusOK,
		Message: "ok",
		User: *NewUserResponse(user),
	}, nil
}
