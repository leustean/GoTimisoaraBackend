package models

import (
	"goTimisoaraBackend/db"
	"html"
	"strings"
)

type Tag struct {
	TagId   uint32 `gorm:"primary_key;auto_increment" json:"tagId"`
	TagName string `gorm:"size:100;not null" json:"tagName"`
}

func (tag *Tag) Prepare() {
	tag.TagId = 0
	tag.TagName = html.EscapeString(strings.TrimSpace(tag.TagName))
}

func (tag *Tag) SaveTag() (*Tag, error) {
	var err error

	database := db.GetDB()
	err = database.Debug().Create(&tag).Error

	if err != nil {
		return &Tag{}, err
	}

	return tag, nil
}

func (tag *Tag) FindAllTags() (*[]Tag, error) {
	var err error
	var tags []Tag
	database := db.GetDB()

	err = database.Debug().Model(&Tag{}).Find(&tags).Error

	if err != nil {
		return &[]Tag{}, err
	}

	return &tags, err
}

func (tag *Tag) UpdateTag() (*Tag, error) {
	database := db.GetDB()

	databaseResult := database.Debug().Model(&Tag{}).Where("tagId = ?", tag.TagId).Take(&Tag{}).UpdateColumns(
		map[string]interface{}{
			"tagName": tag.TagName,
		},
	)

	if databaseResult.Error != nil {
		return &Tag{}, databaseResult.Error
	}

	return tag, nil
}

func (tag *Tag) DeleteTag() error {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&Tag{}).Where("tagId = ?", tag.TagId).Take(&Tag{}).Delete(&Tag{})

	if databaseResult.Error != nil {
		return databaseResult.Error
	}

	return nil
}
