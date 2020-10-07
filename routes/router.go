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
		// user module router
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		// category module router
		router.POST("catagory/add", v1.AddCatagory)
		router.GET("catagories", v1.GetCatagories)
		router.PUT("catagory/:id", v1.EditCatagory)
		router.DELETE("catagory/:id", v1.DeleteCatagory)

		// article module router
	}

	return engine
}
