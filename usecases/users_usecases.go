package usecases

import (
    "financial-manager-api/dtos"
    "financial-manager-api/models"
    "financial-manager-api/pkg/crypto"
    "financial-manager-api/repositories"
    "financial-manager-api/utils/logger"
)

type UsersUsecase struct {
    repository repositories.UserRepository
}

func NewUsersUsecase(repository repositories.UserRepository) UsersUsecase {
    return UsersUsecase{
        repository: repository,
    }
}

func (uu *UsersUsecase) GetUsers() ([]models.UsersModel, error) {
    users, err := uu.repository.GetUsers()
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (uu *UsersUsecase) CreateUser(createUserData dtos.UserRequest) (models.UsersModel, error) {
    hashedPassword, hashPasswordError := crypto.HashPassword(createUserData.Password)
    if hashPasswordError != nil {
        logger.L.Errorw("Erro ao converter a senha do usuário para hash!",
            "userEmail", createUserData.Email,
            "error", hashPasswordError,
        )
        return models.UsersModel{}, hashPasswordError
    }
    createUserData.Password = hashedPassword
    userCreated, createUserError := uu.repository.CreateUser(createUserData)
    if createUserError != nil {
        logger.L.Errorw("Erro ao inserir usuário no banco!",
            "userEmail", createUserData.Email,
            "error", createUserError,
        )
    }
    logger.L.Infow("Usuário criado com sucesso!", "userEmail", createUserData.Email)
    return userCreated, nil
}
