package response

import "time"

type TokenResponse struct {
	Value string    `json:"value"`
	Exp   time.Time `json:"expires_at"`
}

type SignUpResponse struct {
	Message string        `json:"message"`
	Token   TokenResponse `json:"token"`
}

func NewSignUpResponse(token string, exp *time.Time) (*SignUpResponse, error) {
	return &SignUpResponse{
		Message: "User created successfully",
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
	}, nil
}

type SignInResponse struct {
	Message        string        `json:"message"`
	Token          TokenResponse `json:"token"`
	OrganizationID string        `json:"organization_id"`
}

func NewSignInResponse(token string, exp *time.Time, orgID string) (*SignInResponse, error) {
	return &SignInResponse{
		Message: "Signed in successfully",
		Token: TokenResponse{
			Value: token,
			Exp:   *exp,
		},
		OrganizationID: orgID,
	}, nil
}
