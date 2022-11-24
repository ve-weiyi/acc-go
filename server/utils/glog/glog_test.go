package glog

import (
	"fmt"
	"log"
	"os"
)

func SyslogFormat() {
	logFile, err := os.OpenFile("./log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
