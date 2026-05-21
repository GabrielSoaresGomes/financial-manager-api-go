package controllers

import (
	"financial-manager-api/dtos"
	"financial-manager-api/pkg/app_errors"
	"financial-manager-api/usecases"
	"financial-manager-api/utils/logger"
	"net/http"
	"strconv"

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
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.FromUsersModelToResponse(users))
}

func (uc *UsersController) GetUserById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, app_errors.ErrorBadRequest("id inválido"))
		return
	}

	userFound, err := uc.usersUsecase.GetUserById(userId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.FromUserModelToResponse(userFound))
}

func (uc *UsersController) CreateUser(ctx *gin.Context) {
	var createUserData dtos.UserRequest
	if bindBodyError := ctx.ShouldBindJSON(&createUserData); bindBodyError != nil {
		logger.L.Errorw("Campo do JSON inválido", "error", bindBodyError)
		handleError(ctx, app_errors.ErrorBadRequest(bindBodyError.Error()))
		return
	}

	createdUser, err := uc.usersUsecase.CreateUser(createUserData)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dtos.FromUserModelToResponse(createdUser))
}
