package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)

	code = model.CheckCategoryExist(data.Name)
	if code == errmsg.Success {
		model.CreateCategory(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
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

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.CheckCategoryExist(data.Name)
	if code == errmsg.Success {
		model.EditCategory(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
