package models

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"goTimisoaraBackend/db"
	"html"
	"strings"
	"time"
)

type Article struct {
	ArticleID     uint32          `gorm:"primary_key;auto_increment" json:"articleId"`
	Title         string          `gorm:"size:255;not null;unique" json:"title"`
	Author        User            `json:"author"`
	AuthorID      uint32          `gorm:"not null" json:"authorId"`
	Content       json.RawMessage `gorm:"type:JSON; not null" json:"contents"`
	IsVisible     uint            `json:"isVisible"`
	Tag           Tag             `json:"tag"`
	TagId         uint32          `json:"tagId"`
	EditorsChoice uint            `json:"editorsChoice"`
	ViewCount     uint32          `json:"viewCount"`
	UpdatedAt     time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
	CreatedAt     time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt,omitempty"`
}

func (article *Article) Prepare() {
	article.Title = html.EscapeString(strings.TrimSpace(article.Title))
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

	databaseResult := database.Debug().Model(&Article{}).Where("article_id = ?", article.ArticleID).Take(&Article{}).UpdateColumns(
		map[string]interface{}{
			"title":          article.Title,
			"tag_id":         article.TagId,
			"updated_at":     article.UpdatedAt,
			"is_visible":     article.IsVisible,
			"editors_choice": article.EditorsChoice,
			"content":        article.Content,
		},
	)

	if databaseResult.Error != nil {
		return &Article{}, databaseResult.Error
	}

	return article, nil
}

func (article *Article) DeleteArticleById(articleId uint32) (int64, error) {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&Article{}).Where("article_id = ?", articleId).Take(&Article{}).Delete(&Article{})

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

func (article *Article) FindArticlesByPageNumber(pageNumber uint32, tagId uint32, sortType uint8) (uint32, *[]Article, error) {
	var err error
	var articles []Article
	var numberOfResultsOnPage uint32 = 10
	var computedPage uint32 = 1
	var pageCount uint32 = 0

	if pageNumber >= 1 {
		computedPage = (pageNumber - 1) * numberOfResultsOnPage
	}

	database := db.GetDB()

	if tagId != 0 && sortType == 0 {
		err = database.Debug().Model(&Article{}).Offset(computedPage).Limit(numberOfResultsOnPage).Where("tag_id = ?", tagId).Find(&articles).Error
		err = database.Debug().Model(&Article{}).Where("tag_id = ?", tagId).Count(&pageCount).Error
	} else if tagId != 0 && sortType != 0 {
		if sortType == 1 {
			err = database.Debug().Model(&Article{}).Where("tag_id = ?", tagId).Order("view_count desc").Order("updated_at desc").Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
			err = database.Debug().Model(&Article{}).Where("tag_id = ?", tagId).Count(&pageCount).Error
		} else if sortType == 2 {
			err = database.Debug().Model(&Article{}).Where("editors_choice = 1 AND tag_id = ?", tagId).Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
			err = database.Debug().Model(&Article{}).Where("editors_choice = 1 AND tag_id = ?", tagId).Count(&pageCount).Error
		} else {
			err = database.Debug().Model(&Article{}).Where("tag_id = ?", tagId).Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
			err = database.Debug().Model(&Article{}).Where("tag_id = ?", tagId).Count(&pageCount).Error
		}
	} else if tagId == 0 && sortType != 0 {
		if sortType == 1 {
			err = database.Debug().Model(&Article{}).Order("view_count desc").Order("updated_at desc").Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
			err = database.Debug().Model(&Article{}).Count(&pageCount).Error
		} else if sortType == 2 {
			err = database.Debug().Model(&Article{}).Where("editors_choice = 1").Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
			err = database.Debug().Model(&Article{}).Where("editors_choice = 1").Count(&pageCount).Error
		} else {
			err = database.Debug().Model(&Article{}).Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
			err = database.Debug().Model(&Article{}).Count(&pageCount).Error
		}
	} else {
		err = database.Debug().Model(&Article{}).Offset(computedPage).Limit(numberOfResultsOnPage).Find(&articles).Error
		err = database.Debug().Model(&Article{}).Count(&pageCount).Error
	}

	if err != nil {
		return 0, &[]Article{}, err
	}

	if pageCount == 0 {
		pageCount += 1
	}
	pageCount = ((pageCount - 1) / numberOfResultsOnPage) + 1
	return pageCount, &articles, err
}

func (article *Article) FindArticleById(articleId uint32) (*Article, error) {
	var err error
	var articleResult Article
	database := db.GetDB()

	err = database.Debug().Model(&Article{}).Where("article_id = ?", articleId).Limit(1).Find(&articleResult).Error

	if err != nil {
		return &Article{}, err
	}

	err = database.Debug().Model(&User{}).Where("user_id = ?", articleResult.AuthorID).Limit(1).Find(&articleResult.Author).Error

	if err != nil {
		return &Article{}, err
	}

	_ = database.Debug().Model(&Tag{}).Where("tag_id = ?", articleResult.TagId).Limit(1).Find(&articleResult.Tag).Error

	_ = database.Debug().Model(&Article{}).Where("article_id = ?", articleId).Take(&Article{}).UpdateColumns(
		map[string]interface{}{
			"view_count": articleResult.ViewCount + 1,
		},
	)

	return &articleResult, err
}
