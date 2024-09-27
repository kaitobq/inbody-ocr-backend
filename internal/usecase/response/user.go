package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"time"
)

type UserResponse struct {
	ID   string                  `json:"id"`
	Name string                  `json:"name"`
	Role entity.OrganizationRole `json:"role"`
}

type TokenResponse struct {
	Value string    `json:"value"`
	Exp   time.Time `json:"expires_at"`
}

type SignUpResponse struct {
	Message string        `json:"message"`
	Token   TokenResponse `json:"token"`
	User    UserResponse  `json:"user"`
}

func NewSignUpResponse(token string, exp *time.Time, userID, userName string, role entity.OrganizationRole) (*SignUpResponse, error) {
	return &SignUpResponse{
		Message: "User created successfully",
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		User: UserResponse{
			ID:   userID,
			Name: userName,
			Role: role,
		},
	}, nil
}

type SignInResponse struct {
	Message        string        `json:"message"`
	Token          TokenResponse `json:"token"`
	OrganizationID string        `json:"organization_id"`
	User           UserResponse  `json:"user"`
}

func NewSignInResponse(token string, exp *time.Time, orgID, userID, userName string, role entity.OrganizationRole) (*SignInResponse, error) {
	return &SignInResponse{
		Message: "Signed in successfully",
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		OrganizationID: orgID,
		User: UserResponse{
			ID:   userID,
			Name: userName,
			Role: role,
		},
	}, nil
}
