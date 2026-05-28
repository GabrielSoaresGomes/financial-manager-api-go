package controllers

import (
	"errors"
	"financial-manager-api/pkg/app_errors"
	"financial-manager-api/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context, err error) {
	var appError *app_errors.AppError
	if errors.As(err, &appError) {
		if appError.Cause != nil {
			logger.L.Errorw("Erro de aplicação",
				"status", appError.StatusCode,
				"message", appError.Message,
				"cause", appError.Cause.Error(),
			)
		} else {
			logger.L.Errorw("Erro de aplicação",
				"status", appError.StatusCode,
				"message", appError.Message,
			)
		}
		ctx.JSON(appError.StatusCode, gin.H{"error": appError.Message})
		return
	}
	logger.L.Errorw("Erro interno não mapeado",
		"error", err.Error(),
	)
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno no servidor"})
}
