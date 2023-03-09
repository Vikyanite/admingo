package zlog

import "gopkg.in/natefinch/lumberjack.v2"

type Config struct {
	JSONEncoder bool `yaml:"JSONEncoder"`
	// 因为zaplog没有分割日志文件的功能，所以采用lumberjack
	logRoll lumberjack.Logger `yaml:"lumberjack"`
}

var c *Config

func init() {

}
