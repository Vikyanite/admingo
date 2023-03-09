package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

func Info(msg string) {
	do(zap.InfoLevel, "", msg)
}

func Debug(msg string) {
	do(zap.DebugLevel, "", msg)
}

func Warn(msg string) {
	do(zap.WarnLevel, "", msg)
}

func Error(msg string) {
	do(zap.ErrorLevel, "", msg)
}

func Panic(msg string) {
	do(zap.PanicLevel, "", msg)
}

func Fatal(msg string) {
	do(zap.FatalLevel, "", msg)
}

func Infof(tmplt string, args ...interface{}) {
	do(zap.InfoLevel, tmplt, args...)
}

func Debugf(tmplt string, args ...interface{}) {
	do(zap.DebugLevel, tmplt, args...)
}

func Warnf(tmplt string, args ...interface{}) {
	do(zap.WarnLevel, tmplt, args...)
}

func Errorf(tmplt string, args ...interface{}) {
	do(zap.ErrorLevel, tmplt, args...)
}

func Panicf(tmplt string, args ...interface{}) {
	do(zap.PanicLevel, tmplt, args...)
}

func Fatalf(tmplt string, args ...interface{}) {
	do(zap.FatalLevel, tmplt, args...)
}

func do(l zapcore.Level, tmplt string, args ...interface{}) {
	field := getCallerInfoForLog()
	msg := formatMsg(tmplt, args)
	switch l {
	case zapcore.DebugLevel:
		logger.Debug(msg, field...)
	case zapcore.InfoLevel:
		logger.Info(msg, field...)
	case zapcore.WarnLevel:
		logger.Warn(msg, field...)
	case zapcore.ErrorLevel:
		logger.Error(msg, field...)
	case zapcore.PanicLevel:
		logger.Panic(msg, field...)
	case zapcore.FatalLevel:
		logger.Fatal(msg, field...)
	}

}

func formatMsg(template string, args []interface{}) string {
	if template == "" {
		return args[0].(string)
	}
	return fmt.Sprintf(template, args...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名
	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
