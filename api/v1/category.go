package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddCatagory 添加分类
func AddCatagory(c *gin.Context) {
	var data model.Catagory
	c.ShouldBindJSON(&data)

	code = model.CheckCatagoryExist(data.Name)
	if code == errmsg.Success {
		model.CreateCatagory(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetCatagories 查询分类列表
func GetCatagories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCatagories(pageSize, pageNum)
	code = errmsg.Success

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// EditCatagory 编辑分类
func EditCatagory(c *gin.Context) {
	var data model.Catagory
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.CheckCatagoryExist(data.Name)
	if code == errmsg.Success {
		model.EditCatagory(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// DeleteCatagory 删除分类
func DeleteCatagory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCatagory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
