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
	imageDataCtrl *ImageDataController,
	tokenService service.TokenService,
) {
	v1 := r.Group("api/v1")

	user := v1.Group("user")
	user.Use(middleware.AuthMiddleware(tokenService))
	{
		user.GET("", userCtrl.GetOwnInfo)
	}

	organization := v1.Group("organization")
	{
		organization.POST("", organizationCtrl.CreateOrganization)
		organization.POST("/:id/signup", userCtrl.SignUp)
		organization.POST("/signin", userCtrl.SignIn)
	}
	organization.Use(middleware.AuthMiddleware(tokenService))
	{
		organization.GET("/dashboard", organizationCtrl.GetScreenDashboard)
		organization.GET("/dashboard/admin", organizationCtrl.GetScreenDashboardForAdmin)
	}

	image := v1.Group("image")
	image.Use(middleware.AuthMiddleware(tokenService))
	{
		image.POST("", imageCtrl.AnalyzeImage)
	}

	imageData := v1.Group("image-data")
	imageData.Use(middleware.AuthMiddleware(tokenService))
	{
		imageData.POST("", imageDataCtrl.SaveImageData)
		imageData.GET("", imageDataCtrl.GetImageDataForMember)
		imageData.GET("/admin", imageDataCtrl.GetImageDataForAdmin)
	}
}
