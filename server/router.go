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
	}
	return router

}
