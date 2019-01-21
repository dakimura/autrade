package logger

import (
	"go.uber.org/zap"
	"sync"
)

var once sync.Once


func Init() error {
	var err error
	var logger *zap.Logger
	// once ensures the singleton is initialized only once
	once.Do(func() {
		logger, err = zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
	})
	if err != nil {
		return err
	}
	return nil
}

//
//// Log a message at the given level with given fields
//func Log(level zapcore.Level, message string, fields ...zap.Field) {
//	if ce := singleton.Check(level, message); ce != nil {
//		ce.Write(fields...)
//	}
//}
//
//// Debug logs a debug message with the given fields
//func Debug(message string, fields ...zap.Field) {
//	singleton.Debug(message, fields...)
//}
//
//// Info logs a debug message with the given fields
//func Info(message string, fields ...zap.Field) {
//	singleton.Info(message, fields...)
//}
//
//// Warn logs a debug message with the given fields
//func Warn(message string, fields ...zap.Field) {
//	singleton.Warn(message, fields...)
//}
//
//// Error logs a debug message with the given fields
//func Error(message string, fields ...zap.Field) {
//	singleton.Error(message, fields...)
//}
//
//// Fatal logs a message than calls os.Exit(1)
//func Fatal(message string, fields ...zap.Field) {
//	singleton.Fatal(message, fields...)
//}
