package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/forms/user"
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
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user"})
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
	var userLoginForm userForms.UserLogin
	var userData models.User

	err := c.Bind(&userData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	userLoginForm.Username = userData.Username
	userLoginForm.Password = userData.Password

	validationResult := userLoginForm.Validate()

	if validationResult != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	user, err := userModel.FindUserByUsername(userData.Username)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user"})
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
	var userRegisterForm userForms.UserRegister

	err := c.Bind(&userData)
	userRegisterForm.Username = userData.Username
	userRegisterForm.Email = userData.Email
	userRegisterForm.FullName = userData.FullName
	userRegisterForm.Password = userData.Password

	validationResult := userRegisterForm.Validate()

	if err != nil {
		log.Println("invalid request: " + err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	if validationResult != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	_, err = userModel.FindUserByUsername(userData.Username)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User already exists"})
		c.Abort()
		return
	}

	userData.Prepare()
	err = userData.BeforeSave()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something wrong happened"})
		c.Abort()
		return
	}

	_, err = userData.SaveUser()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something wrong happened"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
	c.Abort()
	return
}
