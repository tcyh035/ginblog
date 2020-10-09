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
		router.POST("category/add", v1.AddCategory)
		router.GET("categories", v1.GetCategories)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)

		// article module router
		router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArticles)
		router.GET("article/list/:category", v1.GetArticlesByCategory)
		router.GET("article/info/:id", v1.GetArticle)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)
	}

	return engine
}
