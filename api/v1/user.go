package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	c.ShouldBindJSON(&data)
	msg, code = validator.Validate(data)

	if code != errmsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})

		return
	}

	code = model.CheckUserExist(data.Username)
	if code == errmsg.Success {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.Success

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMessage(code),
	})
}

//EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.CheckUserExist(data.Username)
	if code == errmsg.Success {
		model.EditUser(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
