package request

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func GetImgFileFromContext(c *gin.Context) (multipart.File, error) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return file, nil
}
