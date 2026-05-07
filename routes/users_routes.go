package routes

import (
	"database/sql"
	"financial-manager-api/controllers"
	"financial-manager-api/repositories"
	"financial-manager-api/usecases"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(rg *gin.RouterGroup, dbConnection *sql.DB) {
	userRepository := repositories.NewUserRepository(dbConnection)
	usersUsecase := usecases.NewUsersUsecase(userRepository)
	userController := controllers.NewUserController(usersUsecase)

	usersGroup := rg.Group("/users")
	{
		usersGroup.GET("/", userController.GetUsers)
	}
}
