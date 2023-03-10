package zlog

import (
	"errors"
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
)

func init() {
	logConfig = &Config{
		logRoll: lumberjack.Logger{
			Filename:   "../../log/test.log",
			MaxSize:    100,
			MaxAge:     1,
			MaxBackups: 60,
			Compress:   false,
		},
	}
	Init()
}

func TestInfo(t *testing.T) {
	Info("hi")
	Infof("hi")
	Infof("hihi: %d", 1)
	Infof("hihi")
}

func TestError(t *testing.T) {
	Error("hi")
	Errorf("hi")
	Errorf("hihi: %d", 1)
	Errorf("hihi")
}

func TestDebug(t *testing.T) {
	Debug("hi")
	Debugf("hi")
	Debugf("hihi: %d", 1)
	Debugf("hihi")
}

func TestWarn(t *testing.T) {
	Warn("hi")
	Warnf("hi")
	Warnf("hihi: %d", 1)
	Warnf("hihi")
}

func TestFatalf(t *testing.T) {
	Fatalf("hi")
}

func TestPanicf(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	Panicf("hi")
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	Panic("hi")
}

// fatal的话底层是os.Exit(0)所以不大能捕获...
func TestFatal(t *testing.T) {
	Fatal("hi")
}

func TestMustNil(t *testing.T) {
	MustNil(errors.New("hi"))
}
