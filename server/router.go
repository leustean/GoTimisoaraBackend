package server

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.POST("/auth", user.Authentication)
			userGroup.POST("/register", user.Register)
		}

		tagGroup := v1.Group("tag")
		{
			tag := new(controllers.TagController)
			tagGroup.PUT("/", tag.Put)
			tagGroup.POST("/", tag.Post)
			tagGroup.DELETE("/:id", tag.Delete)
			tagGroup.GET("/", tag.GetAll)
			tagGroup.GET("/:id", tag.GetTagById)
		}

		authorGroup := v1.Group("author")
		{
			author := new(controllers.AuthorController)
			authorGroup.PUT("/", author.Put)
			authorGroup.POST("/", author.Post)
			authorGroup.DELETE("/:id", author.Delete)
			authorGroup.GET("/", author.GetAll)
		}

		articleGroup := v1.Group("article")
		{
			article := new(controllers.ArticleController)
			articleGroup.PUT("/", article.Put)
			articleGroup.POST("/", article.Post)
			articleGroup.DELETE("/:id", article.Delete)
			articleGroup.GET("/", article.GetAll)
		}
	}

	return router
}
