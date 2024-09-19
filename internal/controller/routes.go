package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
	organizationCtrl *OrganizationController,
	membershipCtrl *UserOrganizationMembershipController,
	tokenService service.TokenService,
) {
	v1 := r.Group("api/v1")

	auth := v1.Group("auth")
	{
		auth.POST("/signup", userCtrl.SignUp)
		auth.POST("/signin", userCtrl.SignIn)
	}

	organization := v1.Group("organization")
	organization.Use(middleware.AuthMiddleware(tokenService))
	{
		organization.POST("/", organizationCtrl.CreateOrganization)
	}
}
