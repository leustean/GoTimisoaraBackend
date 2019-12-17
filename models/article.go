package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"goTimisoaraBackend/db"
	"html"
	"strings"
	"time"
)

type Article struct {
	ArticleID     uint32    `gorm:"primary_key;auto_increment" json:"articleId"`
	Title         string    `gorm:"size:255;not null;unique" json:"title"`
	Author        User      `json:"author"`
	AuthorID      uint32    `gorm:"not null" json:"authorId"`
	Content       string    `gorm:"type:longtext; not null" json:"contents"`
	IsVisible     uint      `json:"isVisible,omitempty"`
	Tag           Tag       `json:"tag"`
	TagId         uint32    `json:"tagId"`
	EditorsChoice bool      `json:"editorsChoice"`
	ViewCount     uint32    `json:"viewCount"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt,omitempty"`
}

func (article *Article) Prepare() {
	article.Title = html.EscapeString(strings.TrimSpace(article.Title))
	article.Content = html.EscapeString(strings.TrimSpace(article.Content))
	article.Author = User{}
	article.Tag = Tag{}

	if article.ArticleID > 0 {
		authorData, err := article.Author.FindUserById(article.AuthorID)

		if err == nil {
			article.Author = authorData
		}

		tagData, err := article.Tag.FindTagById(article.TagId)

		if err == nil {
			article.Tag = tagData
		}
	}

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
}

func (article *Article) Validate() error {
	if article.Title == "" {
		return errors.New("required Title")
	}
	if article.Content == "" {
		return errors.New("required Content")
	}
	if article.AuthorID < 1 {
		return errors.New("required Author")
	}

	return nil
}

func (article *Article) SaveArticle() (*Article, error) {
	var err error

	database := db.GetDB()
	err = database.Debug().Create(&article).Error

	if err != nil {
		return &Article{}, err
	}

	return article, nil
}

func (article *Article) UpdateArticle() (*Article, error) {
	database := db.GetDB()

	databaseResult := database.Debug().Model(&Article{}).Where("id = ?", article.ArticleID).Take(&Article{}).UpdateColumns(
		map[string]interface{}{
			"title":          article.Title,
			"author_id":      article.AuthorID,
			"tag_id":         article.TagId,
			"updated_at":     article.UpdatedAt,
			"view_count":     article.ViewCount,
			"editors_choice": article.EditorsChoice,
			"contents":       article.Content,
		},
	)

	if databaseResult.Error != nil {
		return &Article{}, databaseResult.Error
	}

	return article, nil
}

func (article *Article) DeleteArticleById(articleId uint64) (int64, error) {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&Article{}).Where("id = ?", articleId).Take(&Article{}).Delete(&Article{})

	if databaseResult.Error != nil {
		if gorm.IsRecordNotFoundError(databaseResult.Error) {
			return 0, errors.New("article not found")
		}

		return 0, databaseResult.Error
	}

	return databaseResult.RowsAffected, nil
}

func (article *Article) FindAllArticles() (*[]Article, error) {
	var err error
	var articles []Article
	database := db.GetDB()

	err = database.Debug().Model(&Article{}).Find(&articles).Error

	if err != nil {
		return &[]Article{}, err
	}

	return &articles, err
}

func (article *Article) FindArticleById(articleId uint32) (Article, error) {
	var err error
	var articleResult Article
	database := db.GetDB()

	err = database.Debug().Model(&Article{}).Where("id = ?", articleId).Limit(1).Find(&articleResult).Error

	if err != nil {
		return Article{}, err
	}

	return articleResult, err
}
