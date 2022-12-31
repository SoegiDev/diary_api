package model

import (
	"html"
	"strings"

	"diary_api/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:255;not null;unique" json:"username"`
	Password     string `gorm:"size:255;not null;" json:"-"`
	Entries_data []Entry
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	return user, err
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).First(&user).Error
	return user, err
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Entries_data").Where("ID=?", id).First(&user).Error
	return user, err
}
