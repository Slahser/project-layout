package edgeserver

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
	Logger *zap.Logger
)

//InitLogger
func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	Logger = zap.New(core, zap.AddCaller())
	Logger.Named(AppConfig.Log.LoggerName)
	return Logger
}

//getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//getLogWriter
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   AppConfig.Log.OutputFile,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     AppConfig.Log.LogKeepDays,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
