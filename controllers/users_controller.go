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
	var filter dtos.UserFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		handleError(ctx, app_errors.ErrorBadRequest(err.Error()))
		return
	}
	users, err := uc.usersUsecase.GetUsers(filter)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.FromUsersModelToResponse(users))
}

func (uc *UsersController) GetUserById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, app_errors.ErrorBadRequest("id do usuário é inválido"))
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

func (uc *UsersController) UpdateUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.L.Errorw("ID do usuário é inválido", "error", err.Error())
		handleError(ctx, app_errors.ErrorBadRequest("id do usuário é inválido"))
		return
	}

	var updateUserData dtos.UserRequest
	if bindBodyError := ctx.ShouldBindJSON(&updateUserData); bindBodyError != nil {
		logger.L.Errorw("Campo do JSON é inválido", "error", bindBodyError)
		handleError(ctx, app_errors.ErrorBadRequest(bindBodyError.Error()))
		return
	}

	updatedUser, err := uc.usersUsecase.UpdateUser(userId, updateUserData)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.FromUserModelToResponse(updatedUser))
}

func (uc *UsersController) DeleteUserById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.L.Errorw("ID do usuário é inválido", "error", err.Error())
		handleError(ctx, app_errors.ErrorBadRequest("id do usuário é inválido"))
		return
	}

	if userId <= 0 {
		logger.L.Errorf("ID do usuário: %d é menor ou igual a 0", userId)
		handleError(ctx, app_errors.ErrorBadRequest("ID do usuário é inválido"))
		return
	}

	deletedUserId, err := uc.usersUsecase.DeleteUserById(userId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if deletedUserId <= 0 {
		handleError(ctx, app_errors.ErrorBadRequest("ID retornado ao apagar usuário é inválido"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": deletedUserId})
}
