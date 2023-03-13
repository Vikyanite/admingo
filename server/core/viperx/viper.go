package viperx

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	for k, v := range configMp {
		vip := viper.New() //一个viper实例只能读取一个配置
		vip.SetConfigType("yaml")
		vip.AddConfigPath("./config") //设置读取的文件路径
		vip.SetConfigName(k)
		// 读取配置
		if err := vip.ReadInConfig(); err != nil {
			log.Fatal("读取不到配置文件：" + err.Error())
		}
		if err := vip.Unmarshal(v); err != nil {
			log.Fatal("Unmarshal失败！ ---- " + err.Error())
		}
	}
}
