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
	var authorsData models.Author
	var resultData *models.Author

	err := c.Bind(&authorsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorsData.Prepare()
	resultData, err = authorsData.SaveAuthor()

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
	var authorsData models.Author
	var resultData *models.Author

	err := c.Bind(&authorsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorsData.Prepare()
	resultData, err = authorsData.UpdateAuthor()

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
	var authorsData models.Author

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorId, err := strconv.Atoi(c.Param("id"))

	err = authorsData.DeleteAuthorById(uint32(authorId))

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
	var authorModel models.Author

	authorsData, err := authorModel.FindAllAuthors()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, authorsData)

	return
}
