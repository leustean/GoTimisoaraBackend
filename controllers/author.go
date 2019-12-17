package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/models"
	"log"
	"net/http"
	"strconv"
)

type AuthorController struct{}

func (author AuthorController) Put(c *gin.Context) {
	var authorsData models.User
	var resultData *models.User

	err := c.Bind(&authorsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorsData.Prepare()
	resultData, err = authorsData.SaveUser()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resultData)

	return
}

func (author AuthorController) Post(c *gin.Context) {
	var authorsData models.User
	var resultData *models.User

	err := c.Bind(&authorsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorsData.Prepare()
	resultData, err = authorsData.UpdateUser()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resultData)

	return
}

func (author AuthorController) Delete(c *gin.Context) {
	var authorsData models.User

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorId, err := strconv.Atoi(c.Param("id"))

	err = authorsData.DeleteUserById(uint32(authorId))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "author has been deleted"})

	return
}

func (author AuthorController) GetAll(c *gin.Context) {
	var authorModel models.User

	authorsData, err := authorModel.FindAllUsers()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, authorsData)

	return
}
