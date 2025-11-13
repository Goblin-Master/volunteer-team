package test

import (
	"testing"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/global"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	configs.LoadConfig()
	global.Init()
	global.Log.Info("TestLog")
	global.Log.Error("test")
	global.Log.Warn("test")
	global.Log.Debug("test")
	global.Log.Errorf("test:%s", "ds")
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("order created", zap.String("orderId", "123456"))
}
