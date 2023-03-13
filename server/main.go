package main

import (
	"github.com/gin-gonic/gin"
	"server/core/viperx"
	"server/core/zlog"
)

func main() {
	viperx.Init()
	zlog.Init()
	engine := gin.Default()
	zlog.Info("hi")
	engine.Run(":3333")
}
