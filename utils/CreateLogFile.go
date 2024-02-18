package utils

import (
	"log"
	"os"
	"path"
	"time"
)

// 创建 logs 文件用来存放日志文件
// 创建日志文件，以每天的日期作为文件名
func CreateLogFile() *os.File {
	// 创建保存日志的文件夹 logs
	// MkdirAll 方法在目录不存在时就创建目录，当目录存在则啥都不做
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		log.Println("dir logs create failed")
	}
	name := time.Now().Format("20060102") + ".log"
	p := path.Join("./logs/", name)
	// O_CREATE 表示文件不存在时就创建
	// O_APPEND 表示文件存在就追加
	file, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Println("log file create failed")
	}
	return file
}
