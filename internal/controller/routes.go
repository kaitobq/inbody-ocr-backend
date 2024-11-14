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
	measurementDateCtrl *MeasurementDateController,
	middleware *middleware.Middleware,
) {
	r.Use(middleware.CORS.CORS())
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
		member := imageData.Group("member")
		member.Use(middleware.API.GuaranteeMember())
		member.POST("", imageDataCtrl.SaveImageData)
		member.GET("/stats", imageDataCtrl.GetStatsForMember)
		member.GET("/chart", imageDataCtrl.GetChartDataForMember)
		member.GET("/data", imageDataCtrl.GetImageDataForMember)
	}

	{
		admin := imageData.Group("admin")
		admin.Use(middleware.API.GuaranteeAdminOROwner())
		admin.GET("/stats", imageDataCtrl.GetStatsForAdmin)
		admin.GET("/chart", imageDataCtrl.GetChartDataForAdmin)
		admin.GET("/data", imageDataCtrl.GetImageDataForAdmin)
		admin.GET("/data/current", imageDataCtrl.GetCurrentImageDataForAdmin)
	}

	measurementDate := v1.Group("measurement-date")
	measurementDate.Use(middleware.API.VerifyToken())
	{
		measurementDate.GET("", measurementDateCtrl.GetMeasurementDate)
	}
}
