package usecases

import (
	"financial-manager-api/dtos"
	"financial-manager-api/models"
	"financial-manager-api/pkg/app_errors"
	"financial-manager-api/pkg/crypto"
	"financial-manager-api/repositories"
	"financial-manager-api/utils/logger"
	"strconv"
)

type UsersUsecase struct {
	repository repositories.UserRepository
}

func NewUsersUsecase(repository repositories.UserRepository) UsersUsecase {
	return UsersUsecase{
		repository: repository,
	}
}

func (uu *UsersUsecase) GetUsers(filter dtos.UserFilter) ([]models.UsersModel, error) {
	users, err := uu.repository.GetUsers(filter)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *UsersUsecase) GetUserById(userId int) (models.UsersModel, error) {
	userFound, err := uu.repository.GetUserById(userId)
	if err != nil {
		return models.UsersModel{}, err
	}
	return userFound, nil
}

func (uu *UsersUsecase) CreateUser(createUserData dtos.UserRequest) (models.UsersModel, error) {
	hashedPassword, hashPasswordError := crypto.HashPassword(createUserData.Password)
	if hashPasswordError != nil {
		logger.L.Errorw("Erro ao converter a senha do usuário para hash!",
			"userEmail", createUserData.Email,
			"error", hashPasswordError,
		)
		return models.UsersModel{}, app_errors.ErrorInternal(hashPasswordError)
	}
	createUserData.Password = hashedPassword

	userCreated, createUserError := uu.repository.CreateUser(createUserData)
	if createUserError != nil {
		return models.UsersModel{}, createUserError
	}

	logger.L.Infow("Usuário criado com sucesso!", "userEmail", createUserData.Email)
	return userCreated, nil
}

func (uu *UsersUsecase) UpdateUser(userId int, updateUserData dtos.UserRequest) (models.UsersModel, error) {
	hashedPassword, hashPasswordError := crypto.HashPassword(updateUserData.Password)
	if hashPasswordError != nil {
		logger.L.Errorw("Erro ao converter a senha do usuário para hash!",
			"userEmail", updateUserData.Email,
			"error", hashPasswordError,
		)
		return models.UsersModel{}, app_errors.ErrorInternal(hashPasswordError)
	}
	updateUserData.Password = hashedPassword

	updatedUser, err := uu.repository.UpdateUser(userId, updateUserData)
	if err != nil {
		stringUserId := strconv.Itoa(userId)
		logger.L.Errorw("Erro ao editar o usuário de ID "+stringUserId, "error", err.Error())
	}

	logger.L.Infow("Usuário editado com sucesso!", "userEmail", updateUserData.Email)
	return updatedUser, nil
}
