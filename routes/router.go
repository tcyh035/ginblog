package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化Router
func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	engine := gin.Default()

	router := engine.Group("api/v1")
	{
		// 用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块路由接口

		// 文章模块路由接口
	}

	return engine
}
