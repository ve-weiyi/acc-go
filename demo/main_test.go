/*
Go 语言的基础组成有以下几个部分：

	包声明
	引入包
	函数
	变量
	语句 & 表达式
	注释
*/

/*
*
1.包名：
你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
*/
package main

import (
	"acc/lib/logger"
	"acc/lib/utils"
	"fmt"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	Success([]string{"11", "22", "33"}, "s", "d")
}

// Success 成功返回
func Success(data ...interface{}) {
	logger.Debug("--->")
	logger.Debug(len(data))
	str := fmt.Sprint(utils.CheckData(data))
	logger.Debug("--->" + str)
	sstr := fmt.Sprint(data[0])
	logger.Debug(sstr)
}

func TestLog(t *testing.T) {

	Index()
}

func Index() {
	//SetFlags为标准记录器设置输出标志。默认标志位是Ldate、Ltime等。
	//log.Ldate：格式是:2009年1月23日
	//log.Ltime：格式是:01:23:23
	log.SetFlags(log.Ldate | log.Ltime)
	//当发生错误或者查看信息的时候，需要查看日志，
	//默认的日志是不显示行号的，
	//可以通过log.SetFlags函数设置显示行号
	//log.LstdFlags：标准默认的日志信息
	//log.Llongfile：显示日志的文件（绝对路径）和对应行号
	//log.Lshortfile：显示日志的文件（不含路径）和对应行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("输出日志")
}
