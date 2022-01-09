package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func init() {
	fileName := "micro.log"
	syncWrite := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    200,
		MaxBackups: 1,
		LocalTime:  true,
		Compress:   true,
	})
	//编码
	encoder := zap.NewProductionEncoderConfig()
	//时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		syncWrite,
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log.Sugar()
}

func Debug(args ...interface{}) {
	Logger.Debug(args)
}

func Warn(args ...interface{})  {
	Logger.Warn(args)
}

func Info (args ...interface{})  {
	Logger.Info(args)
}

func Fatal(args ...interface{})  {
	Logger.Fatal(args)
}