package controller

import (
	"errors"
	"inbody-ocr-backend/internal/controller/render"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/domain/xerror"
	"inbody-ocr-backend/internal/infra/logging"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	uc           usecase.ImageUsecase
	tokenService service.TokenService
}

func NewImageController(uc usecase.ImageUsecase, tokenService service.TokenService) *ImageController {
	return &ImageController{
		uc:           uc,
		tokenService: tokenService,
	}
}

// AnalyzeImage detects text from an uploaded image using Google Vision API
func (ct *ImageController) AnalyzeImage(c *gin.Context) {
	file, fileHeader, err := request.GetImgFileFromContext(c)
	if err != nil {
		logging.Errorf(c, "AnalyzeImage GetImgFileFromContext %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	user := xcontext.MemberUser(c)

	res, err := ct.uc.AnalyzeImage(file, fileHeader, user)
	if err != nil {
		switch {
		case errors.Is(err, xerror.ErrHEICNotSupported):
			logging.Errorf(c, "AnalyzeImage AnalyzeImage err={%v}", err)
			render.ErrorCodeJSON(c, err.Error(), http.StatusBadRequest, xerror.CodeHEICNotSupported)
			return
		default:
			logging.Errorf(c, "AnalyzeImage AnalyzeImage %v", err)
			render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	render.JSON(c, res)
}
