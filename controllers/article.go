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
	var articleData *[]models.Article

	var pageNumber int
	var tagId int
	var sortType int
	var err error

	if c.Param("pageNumber") == "" {
		pageNumber = 1
	} else {
		pageNumber, err = strconv.Atoi(c.Param("pageNumber"))
	}

	if c.Param("tagId") == "" {
		tagId = 0
	} else {
		tagId, err = strconv.Atoi(c.Param("tagId"))
	}

	if c.Param("sortType") == "" {
		sortType = 0
	} else {
		sortType, err = strconv.Atoi(c.Param("sortType"))
	}

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	var pageCount uint32 = 0
	pageCount, articleData, err = articleModel.FindArticlesByPageNumber(uint32(pageNumber), uint32(tagId), uint8(sortType))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"numberOfPages": pageCount, "pageNumber": pageNumber, "articles": articleData})

	return
}

func (article ArticleController) Get(c *gin.Context) {
	var articleModel models.Article
	var articleData *models.Article

	var err error
	var articleId int

	articleId, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	articleData, err = articleModel.FindArticleById(uint32(articleId))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, articleData)

	return
}
