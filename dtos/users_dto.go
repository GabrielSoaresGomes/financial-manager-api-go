package dtos

import (
    "financial-manager-api/enums"
    "financial-manager-api/models"
    "time"
)

type UserResponse struct {
    ID        int                `json:"id"`
    Name      string             `json:"name"`
    Email     string             `json:"email"`
    Role      enums.UserRoleType `json:"role"`
    CreatedAt time.Time          `json:"created_at"`
    UpdatedAt time.Time          `json:"updated_at"`
}

type UserRequest struct {
    Name     string             `json:"name"`
    Email    string             `json:"email"`
    Role     enums.UserRoleType `json:"role"`
    Password string             `json:"password"`
}

func FromUserModelToResponse(userModel models.UsersModel) UserResponse {
    return UserResponse{
        ID:        userModel.ID,
        Name:      userModel.Name,
        Email:     userModel.Email,
        Role:      userModel.Role,
        CreatedAt: userModel.CreatedAt,
        UpdatedAt: userModel.UpdatedAt,
    }
}

func FromUsersModelToResponse(usersModel []models.UsersModel) []UserResponse {
    usersDtoList := make([]UserResponse, len(usersModel))
    for index, value := range usersModel {
        usersDtoList[index] = FromUserModelToResponse(value)
    }
    return usersDtoList
}
