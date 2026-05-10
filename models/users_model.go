package models

import (
    "financial-manager-api/enums"
    "time"
)

type UsersModel struct {
    ID        int                `db:"id" json:"id"`
    Name      string             `json:"name"`
    Email     string             `json:"email"`
    Password  *string            `json:"password"`
    Role      enums.UserRoleType `json:"role"`
    CreatedAt time.Time          `json:"created_at"`
    UpdatedAt time.Time          `json:"updated_at"`
    DeletedAt *time.Time         `json:"deleted_at,omitempty"`
}

func (um *UsersModel) CreateEmpty() UsersModel {
    var usersModel UsersModel
    return usersModel
}
