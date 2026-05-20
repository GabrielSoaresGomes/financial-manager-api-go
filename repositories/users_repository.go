package repositories

import (
	"database/sql"
	"errors"
	"financial-manager-api/dtos"
	"financial-manager-api/enums"
	"financial-manager-api/models"
	"financial-manager-api/utils/logger"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return UserRepository{
		DB: DB,
	}
}

func (ur *UserRepository) GetUsers() ([]models.UsersModel, error) {
	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE deleted_at IS NULL
	`
	rows, err := ur.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		closeRowsError := rows.Close()
		if closeRowsError != nil {
			logger.L.Warnw("Falha ao fechar conexão do retorno da consulta de usuários", "error", closeRowsError)
		}
	}(rows)

	var usersList []models.UsersModel

	for rows.Next() {
		var userObject models.UsersModel
		var roleStr string
		err := rows.Scan(
			&userObject.ID,
			&userObject.Name,
			&userObject.Email,
			&userObject.Password,
			&roleStr,
			&userObject.CreatedAt,
			&userObject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		switch roleStr {
		case "admin":
			userObject.Role = enums.UserRoleAdmin
		case "user":
			userObject.Role = enums.UserRoleRole
		default:
			return nil, errors.New("Tipo do usuário é inválido: " + roleStr)
		}
		usersList = append(usersList, userObject)
	}

	return usersList, nil
}

func (ur *UserRepository) GetUserById(userId int) (models.UsersModel, error) {
	query := `
		SELECT id, name, email, role, created_at, updated_at
		FROM users
		WHERE id = $1
		AND deleted_at IS NULL
	`

	var userObject models.UsersModel
	var roleStr string

	row := ur.DB.QueryRow(query, userId)
	rowScanError := row.Scan(
		&userObject.ID,
		&userObject.Name,
		&userObject.Email,
		&roleStr,
		&userObject.CreatedAt,
		&userObject.UpdatedAt,
	)
	if rowScanError != nil {
		return models.UsersModel{}, rowScanError
	}

	getRoleError := getRoleByRoleString(roleStr, &userObject)
	if getRoleError != nil {
		return models.UsersModel{}, getRoleError
	}
	return userObject, nil
}

func (ur *UserRepository) CreateUser(createUserData dtos.UserRequest) (models.UsersModel, error) {
	query := `
		INSERT INTO users (name, email, role, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, role, created_at, updated_at
	`

	var userObject models.UsersModel
	var roleStr string
	insertExecError := ur.DB.QueryRow(
		query, createUserData.Name,
		createUserData.Email, createUserData.Role, createUserData.Password,
	).Scan(&userObject.ID, &userObject.Name, &userObject.Email, &roleStr, &userObject.CreatedAt, &userObject.UpdatedAt)

	if insertExecError != nil {
		return models.UsersModel{}, insertExecError
	}

	getRoleError := getRoleByRoleString(roleStr, &userObject)
	if getRoleError != nil {
		return models.UsersModel{}, getRoleError
	}

	return userObject, nil
}

func getRoleByRoleString(roleString string, userObject *models.UsersModel) error {
	var role enums.UserRoleType
	switch roleString {
	case "admin":
		role = enums.UserRoleAdmin
	case "user":
		role = enums.UserRoleRole
	default:
		return errors.New("Tipo do usuário é inválido: " + roleString)
	}
	userObject.Role = role
	return nil
}
