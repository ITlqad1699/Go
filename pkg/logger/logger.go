package logger

import (
	"os"

	"github.com/anonydev/e-commerce-api/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	// debug -> info -> warn -> error -> dpanic -> panic -> fatal
	logLevel := config.Log_level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.ErrorLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.InfoLevel
	default:
		level = zap.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.Log_file_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// Format logger
func getEncoderLog() zapcore.Encoder {
	// 1736170682.0589921 -> 2025-01-06T20:38:02.058+0700
	encodeConfiger := zap.NewProductionEncoderConfig()
	encodeConfiger.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfiger.TimeKey = "time"
	// lvl -> level
	encodeConfiger.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller": "main.log.go:line-number"
	encodeConfiger.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfiger)
}
