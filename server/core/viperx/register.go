package viperx

import (
	"log"
)

var configMp = make(map[string]interface{})

// Register 注册要读取的配置文件名与对应需要Unmarshal的struct
func Register(cFilename string, cStruct interface{}) {
	if _, ok := configMp[cFilename]; ok {
		log.Fatalf("config: [%s] registered twice", cFilename)
		return
	}
	configMp[cFilename] = cStruct
}
