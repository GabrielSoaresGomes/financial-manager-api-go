package controllers

import (
	"financial-manager-api/dtos"
	"financial-manager-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersUsecase usecases.UsersUsecase
}

func NewUserController(usecase usecases.UsersUsecase) UsersController {
	return UsersController{
		usersUsecase: usecase,
	}
}

func (uc *UsersController) GetUsers(ctx *gin.Context) {
	users, err := uc.usersUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersResponse := dtos.FromUsersModelToResponse(users)
	ctx.JSON(http.StatusOK, usersResponse)
	return
}

func (uc *UsersController) CreateUser(ctx *gin.Context) {
	var createUserData dtos.UserRequest
	if bindBodyError := ctx.ShouldBindJSON(&createUserData); bindBodyError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": bindBodyError.Error()})
		return
	}

	createdUser, createUserError := uc.usersUsecase.CreateUser(createUserData)
	if createUserError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": createUserError.Error()})
		return
	}

	userResponse := dtos.FromUserModelToResponse(createdUser)
	ctx.JSON(http.StatusCreated, userResponse)
	return
}
