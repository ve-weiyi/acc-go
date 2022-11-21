package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestViper(t *testing.T) {
	Viper("")
}

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path string) *viper.Viper {

	// 1. 初始化 Viper 库
	Viper := viper.New()
	// 2. 环境变量配置文件查找的路径，相对于 当前文件运行目录
	//Viper.AddConfigPath(".")
	//Viper.AddConfigPath("./config") // 添加多个搜索目录
	// 3. 设置文件名称，不带后缀
	Viper.SetConfigName("config")
	// 4. 配置类型，支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	Viper.SetConfigType("yaml")

	// 5. 开始读根目录下的 .env 文件，读不到会报错
	err := Viper.ReadInConfig()

	if err == nil {
		log.Printf("use config file -> %s\n", Viper.ConfigFileUsed())
	} else {
		log.Println(err)
	}

	global := Server{}
	if err = Viper.Unmarshal(&global); err != nil {
		fmt.Println(err)
	}

	jj, _ := json.MarshalIndent(global, "", "\t")
	log.Println(string(jj))

	return Viper
}
