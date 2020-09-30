package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code int

// UserExist 查询用户是否存在
func UserExist() {
	//
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code = model.CheckUserExist(data.Username)
	if code == errmsg.Success {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	// todo
	c.JSON(errmsg.Success, gin.H{
		"name": "yu",
	})
}

//EditUser 编辑用户
func EditUser(c *gin.Context) {

}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}