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
	imageCtrl *ImageController,
	tokenService service.TokenService,
) {
	v1 := r.Group("api/v1")

	organization := v1.Group("organization")
	{
		organization.POST("", organizationCtrl.CreateOrganization)
		organization.POST("/:id/signup", userCtrl.SignUp)
		organization.POST("/signin", userCtrl.SignIn)
	}

	image := v1.Group("image")
	image.Use(middleware.AuthMiddleware(tokenService))
	{
		image.POST("", imageCtrl.AnalyzeImage)
	}
}
