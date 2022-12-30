package model

import (
	"diary_api/database"
	"strconv"

	"gorm.io/gorm"
)

type UpdateContent struct {
	Content string `json:"content"`
	UserID  uint
}

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (Entry) TableName() string {
	return "entries_data"
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func (update_data *Entry) ChangeData(id string, ud UpdateContent) (*Entry, error) {
	i, _ := strconv.ParseInt(id, 10, 64)
	err := database.Database.Model(Entry{}).Where("id = ?", i).Updates(ud).Error
	if err != nil {
		return &Entry{}, err
	}
	res, _ := FindEntryById(id)
	return res, nil
}

func FindEntryById(id string) (*Entry, error) {
	var e Entry
	i, _ := strconv.ParseInt(id, 10, 64)
	err := database.Database.Where("id=?", i).Find(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}
