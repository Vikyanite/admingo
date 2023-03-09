package zlog

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"server/core/viperx"
)

type Config struct {
	// 因为zaplog没有分割日志文件的功能，所以采用lumberjack
	logRoll lumberjack.Logger `yaml:"lumberjack"`
}

var logConfig *Config

func init() {
	viperx.Register("zlog", logConfig)
}
