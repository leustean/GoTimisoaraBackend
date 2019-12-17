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
	ID            uint32    `gorm:"primary_key;auto_increment" json:"articleId"`
	Title         string    `gorm:"size:255;not null;unique" json:"title"`
	Author        User      `json:"author"`
	AuthorID      uint32    `gorm:"not null" json:"authorId"`
	Content       string    `gorm:"type:longtext; not null" json:"contents"`
	IsVisible     uint      `json:"isVisible,omitempty"`
	Tag           Tag       `json:"tag"`
	TagId         uint32    `json:"tagId"`
	EditorsChoice bool      `json:"editorsChoice"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt,omitempty"`
}

func (p *Article) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Article) Validate() error {
	if p.Title == "" {
		return errors.New("required Title")
	}
	if p.Content == "" {
		return errors.New("required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("required Author")
	}
	return nil
}

func (p *Article) SavePost() (*Article, error) {
	database := db.GetDB()

	var err error
	err = database.Debug().Model(&Article{}).Create(&p).Error

	if err != nil {
		return &Article{}, err
	}

	if p.ID != 0 {
		err = database.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Article{}, err
		}
	}

	return p, nil
}

func (p *Article) FindAllPosts() (*[]Article, error) {
	var err error
	var posts []Article
	database := db.GetDB()

	err = database.Debug().Model(&Article{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Article{}, err
	}
	if len(posts) > 0 {
		for i := range posts {
			err := database.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Article{}, err
			}
		}
	}

	return &posts, nil
}

func (p *Article) FindPostByID(pid uint64) (*Article, error) {
	var err error
	database := db.GetDB()
	err = database.Debug().Model(&Article{}).Where("id = ?", pid).Take(&p).Error

	if err != nil {
		return &Article{}, err
	}

	if p.ID != 0 {
		err = database.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error

		if err != nil {
			return &Article{}, err
		}
	}
	return p, nil
}

func (p *Article) UpdateAPost() (*Article, error) {

	var err error
	database := db.GetDB()

	err = database.Debug().Model(&Article{}).Where("id = ?", p.ID).Updates(Article{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error

	if err != nil {
		return &Article{}, err
	}

	if p.ID != 0 {
		err = database.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Article{}, err
		}
	}

	return p, nil
}

func (p *Article) DeleteAPost(pid uint64, uid uint32) (int64, error) {

	database := db.GetDB()
	databaseResult := database.Debug().Model(&Article{}).Where("id = ? and author_id = ?", pid, uid).Take(&Article{}).Delete(&Article{})

	if databaseResult.Error != nil {
		if gorm.IsRecordNotFoundError(databaseResult.Error) {
			return 0, errors.New("post not found")
		}
		return 0, databaseResult.Error
	}

	return databaseResult.RowsAffected, nil
}
