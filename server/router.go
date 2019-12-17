package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(cors.Default())

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

			pageNumberGroup := articleGroup.Group("pageNumber")
			{
				pageNumberGroup.GET("/:pageNumber", article.GetAll)
				pageNumberGroup.GET("/:pageNumber/tagId/:tagId", article.GetAll)
				pageNumberGroup.GET("/:pageNumber/tagId/:tagId/sortType/:sortType", article.GetAll)
				pageNumberGroup.GET("/:pageNumber/sortType/:sortType", article.GetAll)
			}

			tagIdGroup := articleGroup.Group("tagId")
			{
				tagIdGroup.GET("/:tagId", article.GetAll)
				tagIdGroup.GET("/:tagId/sortType/:sortType", article.GetAll)
			}

			articleGroup.GET("/sortType/:sortType", article.GetAll)

		}
	}

	return router
}
