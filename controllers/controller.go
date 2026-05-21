package controllers

import (
	"errors"
	"financial-manager-api/pkg/app_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context, err error) {
	var appError *app_errors.AppError
	if errors.As(err, &appError) {
		ctx.JSON(appError.StatusCode, gin.H{"error": appError.Message})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno no servidor"})
}
