package model

import (
	"diary_api/database"
	"diary_api/schema"
)

type Role schema.Role

func (role *Role) Save() (*Role, error) {
	err := database.Database.Create(&role).Error
	return role, err
}
