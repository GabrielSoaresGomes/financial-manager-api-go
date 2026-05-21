package repositories

import (
	"database/sql"
	"errors"
	"financial-manager-api/dtos"
	"financial-manager-api/enums"
	"financial-manager-api/models"
	"financial-manager-api/pkg/app_errors"
	"financial-manager-api/utils/logger"
	"fmt"

	"github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return UserRepository{
		DB: DB,
	}
}

func (ur *UserRepository) GetUsers(filter dtos.UserFilter) ([]models.UsersModel, error) {
	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE deleted_at IS NULL
	`

	var args []any
	index := 1

	if filter.Name != nil {
		query += fmt.Sprintf(" AND name ILIKE $%d", index)
		args = append(args, "%"+*filter.Name+"%")
		index++
	}

	if filter.Email != nil {
		query += fmt.Sprintf(" AND email ILIKE $%d", index)
		args = append(args, "%"+*filter.Email+"%")
		index++
	}

	if filter.Role != nil {
		query += fmt.Sprintf(" AND role = $%d", index)
		args = append(args, *filter.Role)
		index++
	}

	rows, err := ur.DB.Query(query, args...)
	if err != nil {
		return nil, app_errors.ErrorInternal(err)
	}
	defer func(rows *sql.Rows) {
		closeRowsError := rows.Close()
		if closeRowsError != nil {
			logger.L.Warnw("Falha ao fechar conexão do retorno da consulta de usuários", "error", closeRowsError)
		}
	}(rows)

	usersList := make([]models.UsersModel, 0)

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
			return nil, app_errors.ErrorInternal(err)
		}
		if err := getRoleByRoleString(roleStr, &userObject); err != nil {
			return nil, err
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
		if errors.Is(rowScanError, sql.ErrNoRows) {
			return models.UsersModel{}, app_errors.ErrorNotFound("usuário não encontrado")
		}
		return models.UsersModel{}, app_errors.ErrorInternal(rowScanError)
	}

	if err := getRoleByRoleString(roleStr, &userObject); err != nil {
		return models.UsersModel{}, err
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
		var pqErr *pq.Error
		if errors.As(insertExecError, &pqErr) && pqErr.Code == "23505" {
			return models.UsersModel{}, app_errors.ErrorConflict("email já cadastrado")
		}
		return models.UsersModel{}, app_errors.ErrorInternal(insertExecError)
	}

	if err := getRoleByRoleString(roleStr, &userObject); err != nil {
		return models.UsersModel{}, err
	}

	return userObject, nil
}

func getRoleByRoleString(roleString string, userObject *models.UsersModel) error {
	switch roleString {
	case "admin":
		userObject.Role = enums.UserRoleAdmin
	case "user":
		userObject.Role = enums.UserRoleRole
	default:
		return app_errors.ErrorInternal(errors.New("tipo de usuário inválido: " + roleString))
	}
	return nil
}
