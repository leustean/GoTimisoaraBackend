package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/models"
	"log"
	"net/http"
	"strconv"
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

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	tagId, err := strconv.Atoi(c.Param("id"))

	err = tagsData.DeleteTagById(uint32(tagId))

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

func (tag TagController) GetTagById(c *gin.Context) {
	var tagModel models.Tag

	tagId, err := strconv.Atoi(c.Param("id"))
	tagsData, err := tagModel.FindTagById(uint32(tagId))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, tagsData)

	return
}
