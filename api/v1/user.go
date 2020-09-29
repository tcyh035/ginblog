package v1

import (
	"ginblog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

// UserExist 查询用户是否存在
func UserExist() {
	//
}

// AddUser 添加用户
func AddUser(c *gin.Context) {

}

// 查询单个用户

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
