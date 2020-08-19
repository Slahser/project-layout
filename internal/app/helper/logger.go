package helper

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
  usage:
  global zap logger with following configuration
  zap.S() means SugarLogger
  zap.L() means ZapLogger
*/
var (
	logFile     = "./logs/tt.log"
	logLevel    = zapcore.DebugLevel
	logKeepDays = 30
)

func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     logKeepDays,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
