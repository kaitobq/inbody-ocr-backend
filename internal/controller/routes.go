package controller

import (
	"inbody-ocr-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
	organizationCtrl *OrganizationController,
	imageCtrl *ImageController,
	imageDataCtrl *ImageDataController,
	middleware *middleware.Middleware,
) {
	r.Use(middleware.UserAgent.UserAgent())
	r.Use(middleware.Logging.LoggingRequestTraceID())

	v1 := r.Group("api/v1")

	user := v1.Group("user")
	user.Use(middleware.API.VerifyToken())
	{
		user.GET("", userCtrl.GetOwnInfo)
	}

	organization := v1.Group("organization")
	{
		organization.POST("", organizationCtrl.CreateOrganization)
		organization.POST("/:id/signup", userCtrl.SignUp)
		organization.POST("/signin", userCtrl.SignIn)
	}
	organization.Use(middleware.API.GuaranteeAdminOROwner())
	{
		organization.GET("/role", organizationCtrl.GetAllMembers) // memberも取得できるようにするusecaseが出てくるかも
		organization.PUT("/role", organizationCtrl.UpdateRole)
		organization.DELETE("/role", organizationCtrl.DeleteMember)
	}

	image := v1.Group("image")
	image.Use(middleware.API.GuaranteeMember())
	{
		image.POST("", imageCtrl.AnalyzeImage)
	}

	imageData := v1.Group("image-data")

	{
		member := imageData
		member.Use(middleware.API.GuaranteeMember())
		imageData.POST("", imageDataCtrl.SaveImageData)
		imageData.GET("/stats/member", imageDataCtrl.GetStatsForMember)
		imageData.GET("/chart/member", imageDataCtrl.GetChartDataForMember)
		imageData.GET("/data/member", imageDataCtrl.GetImageDataForMember)
	}

	{
		admin := imageData
		admin.Use(middleware.API.GuaranteeAdminOROwner())
		imageData.GET("/stats/admin", imageDataCtrl.GetStatsForAdmin)
		imageData.GET("/chart/admin", imageDataCtrl.GetChartDataForAdmin)
		imageData.GET("/data/admin", imageDataCtrl.GetImageDataForAdmin)
		imageData.GET("/data/admin/current", imageDataCtrl.GetCurrentImageDataForAdmin)
	}
}
