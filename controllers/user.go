package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/forms"
	"goTimisoaraBackend/models"
	"log"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.FindUserById(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u UserController) Authentication(c *gin.Context) {
	var userData forms.UserLogin

	err := c.Bind(&userData)

	if len(userData.Username) == 0 || len(userData.Password) == 0 {
		log.Println("invalid request")

		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": "Request should contain a username and password fields"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		c.Abort()
		return
	}

	user, err := userModel.FindUserByUsername(userData.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err.Error()})
		c.Abort()
		return
	}

	if userModel.VerifyPassword(user.Password, userData.Password) != nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid request", "error": "Invalid credentials"})
	return
}

func (u UserController) Register(c *gin.Context) {
	var userData models.User
	err := c.Bind(&userData)
	userDataValidate := userData.Validate("update")

	if userDataValidate != nil {
		log.Println("invalid request")

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid request", "error": userDataValidate.Error()})
		c.Abort()
		return
	}

	if err != nil {
		log.Println("invalid request: " + err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid request", "error": err.Error()})
		c.Abort()
		return
	}

	if c.Param("id") != "" {
		user, err := userModel.FindUserById(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err.Error()})
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
