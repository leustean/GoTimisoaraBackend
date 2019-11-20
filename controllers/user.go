package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/models"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.FindUserById(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": user})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
