package initialize

import (
	"acc/server/config"
	"acc/server/global"
	"github.com/spf13/viper"
	"log"
)

// ReadConfig //
// 优先级: 命令行 > 环境变量 > 默认值
func ReadConfig(path string) {

	// 1. 初始化 v 库
	v := viper.New()
	// 2. 环境变量配置文件查找的路径，相对于 当前文件运行目录
	v.AddConfigPath("./")
	//v.AddConfigPath("./config") // 添加多个搜索目录
	// 3. 设置文件名称，不带后缀
	v.SetConfigName(path)
	// 4. 配置类型，支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	v.SetConfigType("yaml")

	// 5. 开始读根目录下的 .env 文件，读不到会报错
	err := v.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("use config file -> %con\n", v.ConfigFileUsed())

	con := config.Server{}
	if err := v.Unmarshal(&con); err != nil {
		log.Println(err)
	}

	//转换json打印
	//jj, err := json.MarshalIndent(con, "", "\t")
	//log.Println(string(jj))

	global.GVA_VP = v
	global.GVA_CONFIG = con
}
