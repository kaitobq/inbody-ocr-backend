package request

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func GetImgFileFromContext(c *gin.Context) (multipart.File, *multipart.FileHeader, error) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		return nil, nil, err
	}
	return file, header, nil
}
