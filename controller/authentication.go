package controller

import (
	"diary_api/helper"
	"diary_api/model"
	"diary_api/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterSchema schema.Register
type LoginSchema schema.Login

func Register(context *gin.Context) {
	var input RegisterSchema

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)

	user := model.User{
		Id:       uuid.New(),
		Roles:    []*schema.Role{{Name: "Admin"}},
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input LoginSchema

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := model.FindUserByUsername(input.Username)
	fmt.Println(&user, err)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong Password"})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt})
}
