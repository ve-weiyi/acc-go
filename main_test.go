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
