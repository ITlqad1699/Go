package initialize

import (
	"github.com/anonydev/e-commerce-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Configuration loaded successfully", zap.String("success", "ok"))
	InitMySQL()
	InitRedis()
	r := InitRouter()
	// listen and serve on 0.0.0.0:8002 (for windows "localhost:8002")
	// r.Run(":8002")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run()
	r.Run(":8002")
}
