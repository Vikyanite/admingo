package viperx

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"testing"
)

func TestInit(t *testing.T) {
	type Person struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
		Sex  bool   `yaml:"sex"`
	}
	test := &Person{}
	Register("test", test)

	// Init 因为 单测的运行路径与main运行不同，所以不能直接调用Init方法
	for k, v := range configMp {
		vip := viper.New()
		vip.SetConfigType("yaml")
		vip.AddConfigPath("../../config") //设置读取的文件路径
		vip.SetConfigName(k)
		if err := vip.ReadInConfig(); err != nil {
			panic("读取不到配置文件：" + err.Error())
		}
		if err := vip.Unmarshal(v); err != nil {
			fmt.Printf("%v", err)
		}
	}

	assert.Equal(t, test, &Person{
		Name: "hi",
		Age:  10,
		Sex:  false,
	})
	//zlog.Debugf("%v", test)
}
