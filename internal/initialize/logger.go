package initialize

import (
	"github.com/anonydev/e-commerce-api/global"
	"github.com/anonydev/e-commerce-api/pkg/logger"
)

func InitLogger() {
	// Initialize logger
	global.Logger = logger.NewLogger(global.Config.Logger)
}
