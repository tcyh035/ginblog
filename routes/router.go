package routes

import (
	"ginblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	engine := gin.Default()

	router := engine.Group("api/v1")
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	return engine
}
