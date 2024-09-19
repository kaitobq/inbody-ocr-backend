package controller

import (
	"inbody-ocr-backend/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
	tokenService service.TokenService,
) {
	v1 := r.Group("api/v1")

	auth := v1.Group("auth")
	{
		auth.POST("/signup", userCtrl.SignUp)
		auth.POST("/signin", userCtrl.SignIn)
	}
}
