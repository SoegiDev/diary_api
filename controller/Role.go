package controller

import (
	"diary_api/model"
	"diary_api/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputRole schema.InputRole

func AddRole(context *gin.Context) {
	var input InputRole

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := model.Role{
		Id:   uuid.New(),
		Name: input.Name,
	}

	savedRole, err := role.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"role": savedRole})
}
