package config

import "go.uber.org/zap"

func GenerateLogger() *zap.Logger {
	// TODO ログ設定yaml読み込ませる
	logger, _ := zap.NewDevelopment()
	return logger
}
