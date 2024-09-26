package request

import "github.com/gin-gonic/gin"

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	OrgID    string
}

func NewSignUpRequest(c *gin.Context) (*SignUpRequest, error) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	req.OrgID = c.Param("id")

	return &req, nil
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	OrgID    string
}

func NewSignInRequest(c *gin.Context) (*SignInRequest, error) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	req.OrgID = c.Param("id")

	return &req, nil
}
