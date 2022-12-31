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
	return entry, err
}

func (update_data *Entry) ChangeData(id string, ud UpdateContent) (Entry, error) {
	i, _ := strconv.ParseInt(id, 10, 64)
	err := database.Database.Model(Entry{}).Where("id = ?", i).Updates(ud).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := FindEntryById(id)
	return res, nil
}

func FindEntryById(id string) (Entry, error) {
	var entry Entry
	i, _ := strconv.ParseInt(id, 10, 64)
	err := database.Database.Where("id=?", i).First(&entry).Error
	return entry, err
}
