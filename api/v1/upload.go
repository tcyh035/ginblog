package v1

import (
	"ginblog/upload"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := upload.File(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"url":     url,
	})
}
