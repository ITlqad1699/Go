package global

import (
	"github.com/anonydev/e-commerce-api/pkg/logger"
	"github.com/anonydev/e-commerce-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
