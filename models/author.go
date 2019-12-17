package models

import (
	"goTimisoaraBackend/db"
	"html"
	"strings"
)

type Author struct {
	UserId   uint32 `binding:"required" gorm:"primary_key;auto_increment" json:"userId"`
	Username string `gorm:"size:255;not null" json:"username"`
	Email    string `gorm:"size:255;not null" json:"email"`
	FullName string `gorm:"size:255;not null" json:"fullName"`
}

func (author *Author) Prepare() {
	author.Username = html.EscapeString(strings.TrimSpace(author.Username))
	author.Email = html.EscapeString(strings.TrimSpace(author.Email))
	author.FullName = html.EscapeString(strings.TrimSpace(author.FullName))
}

func (author *Author) SaveAuthor() (*Author, error) {
	var err error

	database := db.GetDB()
	err = database.Debug().Create(&author).Error

	if err != nil {
		return &Author{}, err
	}

	return author, nil
}

func (author *Author) FindAllAuthors() (*[]Author, error) {
	var err error
	var authors []Author
	database := db.GetDB()

	err = database.Debug().Model(&Author{}).Find(&authors).Error

	if err != nil {
		return &[]Author{}, err
	}

	return &authors, err
}

func (author *Author) UpdateAuthor() (*Author, error) {
	database := db.GetDB()

	databaseResult := database.Debug().Model(&Author{}).Where("user_id = ?", author.UserId).Take(&Author{}).UpdateColumns(
		map[string]interface{}{
			"username":  author.Username,
			"email":     author.Email,
			"full_name": author.FullName,
		},
	)

	if databaseResult.Error != nil {
		return &Author{}, databaseResult.Error
	}

	return author, nil
}

func (author *Author) DeleteAuthorById(id uint32) error {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&Author{}).Where("user_id = ?", id).Take(&Author{}).Delete(&Author{})

	if databaseResult.Error != nil {
		return databaseResult.Error
	}

	return nil
}
