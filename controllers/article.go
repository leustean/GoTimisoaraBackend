package controllers

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/models"
	"log"
	"net/http"
	"strconv"
)

type ArticleController struct{}

func (article ArticleController) Put(c *gin.Context) {
	var articleData models.Article
	var resultData *models.Article

	err := c.Bind(&articleData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	articleData.Prepare()
	resultData, err = articleData.SaveArticle()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resultData)

	return
}

func (article ArticleController) Post(c *gin.Context) {
	var articleData models.Article
	var resultData *models.Article

	err := c.Bind(&articleData)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	articleData.Prepare()
	resultData, err = articleData.UpdateArticle()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resultData)

	return
}

func (article ArticleController) Delete(c *gin.Context) {
	var articleData models.Article

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	authorId, err := strconv.Atoi(c.Param("id"))

	_, err = articleData.DeleteArticleById(uint32(authorId))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article has been deleted"})
	return
}

func (article ArticleController) GetAll(c *gin.Context) {
	var articleModel models.Article

	articleData, err := articleModel.FindAllArticles()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, articleData)

	return
}
