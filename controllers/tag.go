package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/models"
	"log"
	"net/http"
)

type TagController struct{}

func (tag TagController) Put(c *gin.Context) {
	var tagsData models.Tag
	var resultData *models.Tag

	err := c.Bind(&tagsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	tagsData.Prepare()
	resultData, err = tagsData.SaveTag()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resultData)

	return
}

func (tag TagController) Post(c *gin.Context) {
	var tagsData models.Tag
	var resultData *models.Tag

	err := c.Bind(&tagsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	tagsData.Prepare()
	resultData, err = tagsData.UpdateTag()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resultData)

	return
}

func (tag TagController) Delete(c *gin.Context) {
	var tagsData models.Tag

	err := c.Bind(&tagsData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	err = tagsData.DeleteTag()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag has been deleted"})

	return
}

func (tag TagController) GetAll(c *gin.Context) {
	var tagModel models.Tag

	tagsData, err := tagModel.FindAllTags()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, tagsData)

	return
}
