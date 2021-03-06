package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)

	code = model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetArticlesByCategory 查询分类下的所有文章
func GetArticlesByCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	category, _ := strconv.Atoi(c.Param("category"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetArticlesByCategory(category, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetArticle 查询单个文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// GetArticles 查询文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetArticles(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMessage(code),
	})
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
