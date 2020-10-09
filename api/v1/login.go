package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登陆
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.Username, data.Password)

	if code == errmsg.Success {
		token, _ := middleware.GenerateToken(data.Username)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMessage(code),
			"token":   token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMessage(code),
		})
	}
}
