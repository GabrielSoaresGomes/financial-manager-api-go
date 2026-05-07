package logger

import (
	"financial-manager-api/enums"

	"go.uber.org/zap"
)

var L *zap.SugaredLogger

func Init(environment enums.EnvironmentName) {
	var zapLogger *zap.Logger
	var err error

	if environment == enums.EnvironmentProduction {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	L = zapLogger.Sugar()
}

func Sync() {
	if L != nil {
		_ = L.Sync()
	}
}
