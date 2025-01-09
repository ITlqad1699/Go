package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Sugar is a logger that wraps the Logger to provide a more ergonomic, but slightly slower, API.
	// sugar := zap.NewExample().Sugar()
	// sugar.Info("Hello name: %s, age: %d", "@nonydev", 25)
	// // Logger is a fast, structured, leveled logger.
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "@nonydev"), zap.Int("age", 25))

	// logger := zap.NewExample()
	// logger.Info("hello")

	// // Development and Production are two pre-configured loggers that are optimized for different environments.
	// logger, _ = zap.NewDevelopment()
	// logger.Info("hello new Developement")

	// logger, _ = zap.NewProduction()
	// logger.Info("hello new Production")

	// 3. Custom Configuration
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	// This syntax is a shorthand for openning a file with the default permissions.
	// name: file name
	// flag: os.O_CREATE|os.O_APPEND|os.O_WRONLY|...
	// perm: 0666|0755
	os.MkdirAll("log", os.ModePerm)
	file, _ := os.OpenFile("log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
