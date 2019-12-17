package models

import (
	"goTimisoaraBackend/db"
	"html"
	"strings"
)

type Tag struct {
	TagId   uint32 `binding:"required" gorm:"primary_key;auto_increment" json:"tagId"`
	TagName string `gorm:"size:100;not null" json:"tagName"`
}

func (tag *Tag) Prepare() {
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

func (tag *Tag) FindTagById(tagId uint32) (Tag, error) {
	var err error
	var tagResult Tag
	database := db.GetDB()

	err = database.Debug().Model(&Tag{}).Where("tag_id = ?", tagId).Limit(1).Find(&tagResult).Error

	if err != nil {
		return Tag{}, err
	}

	return tagResult, err
}

func (tag *Tag) UpdateTag() (*Tag, error) {
	database := db.GetDB()

	databaseResult := database.Debug().Model(&Tag{}).Where("tag_id = ?", tag.TagId).Take(&Tag{}).UpdateColumns(
		map[string]interface{}{
			"tag_name": tag.TagName,
		},
	)

	if databaseResult.Error != nil {
		return &Tag{}, databaseResult.Error
	}

	return tag, nil
}

func (tag *Tag) DeleteTagById(id uint32) error {
	database := db.GetDB()
	databaseResult := database.Debug().Model(&Tag{}).Where("tag_id = ?", id).Take(&Tag{}).Delete(&Tag{})

	if databaseResult.Error != nil {
		return databaseResult.Error
	}

	return nil
}
