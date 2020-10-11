package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化Router
func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	auth := engine.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// user module router
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// category module router
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// article module router
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// upload file
		auth.POST("upload", v1.Upload)
	}

	router := engine.Group("api/v1")
	router.POST("user/add", v1.AddUser)
	router.GET("users", v1.GetUsers)
	router.GET("categories", v1.GetCategories)
	router.GET("articles", v1.GetArticles)
	router.GET("article/list/:category", v1.GetArticlesByCategory)
	router.GET("article/info/:id", v1.GetArticle)
	router.POST("login", v1.Login)

	return engine
}
