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
	r.Use(middleware.UserAgent())
	r.Use(middleware.LoggingRequestTraceID())

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
		organization.GET("/role", organizationCtrl.GetAllMembers)
		organization.PUT("/role", organizationCtrl.UpdateRole)
		organization.DELETE("/role", organizationCtrl.DeleteMember)
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
		imageData.GET("/stats/member", imageDataCtrl.GetStatsForMember)
		imageData.GET("/stats/admin", imageDataCtrl.GetStatsForAdmin)
		imageData.GET("/chart/member", imageDataCtrl.GetChartDataForMember)
		imageData.GET("/chart/admin", imageDataCtrl.GetChartDataForAdmin)
		imageData.GET("/data/member", imageDataCtrl.GetImageDataForMember)
		imageData.GET("/data/admin", imageDataCtrl.GetImageDataForAdmin)
		imageData.GET("/data/admin/current", imageDataCtrl.GetCurrentImageDataForAdmin)
	}
}
