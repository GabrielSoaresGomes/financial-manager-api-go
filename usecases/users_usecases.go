package usecases

import (
	"financial-manager-api/models"
	"financial-manager-api/repositories"
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
