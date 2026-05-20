package main

import (
	"financial-manager-api/configs/db"
	"financial-manager-api/enums"
	"financial-manager-api/routes"
	"financial-manager-api/utils/logger"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if envLoadError := godotenv.Load(".env"); envLoadError != nil {
		fmt.Println("Aviso: .env não encontrado, usando variáveis do ambiente")
	}

	environmentString := os.Getenv("APP_ENVIRONMENT")
	environment := enums.EnvironmentName(environmentString)
	logger.Init(environment)

	logger.L.Infow("Iniciando conexão com o banco de dados")
	dbConnection, dbConnectError := db.ConnectDB()
	if dbConnectError != nil {
		logger.L.Errorw("Erro ao realizar conexão com o banco", "error", dbConnectError)
		panic(dbConnectError)
	}

	defer func() {
		if dbConnectionCloseError := dbConnection.Close(); dbConnectionCloseError != nil {
			logger.L.Errorw("Erro ao fechar conexão com o banco", "error", dbConnectionCloseError)
		}
	}()

	logger.L.Info("Iniciando execução das migrations")
	if runMigrationsError := db.RunMigrations(); runMigrationsError != nil {
		logger.L.Fatalw("Falha nas migrations", "error", runMigrationsError)
	}
	logger.L.Info("Migrations aplicadas com sucesso")

	logger.L.Infow("Iniciando criação e registro de rotas da API")
	server := gin.Default()
	routes.RegisterRoutes(server, dbConnection)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "2004"
	}
	serverAddress := fmt.Sprintf(":%s", port)
	logger.L.Infof("Iniciando execução do servidor no endereço: %s", serverAddress)
	if serverRunError := server.Run(serverAddress); serverRunError != nil {
		logger.L.Errorw("Erro ao rodar o servidor", "error", serverRunError)
		panic(serverRunError)
	}
}
