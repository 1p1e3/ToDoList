package main

import (
	"ToDoList/models"
	"ToDoList/routers"
	"ToDoList/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)

func init() {
	// 日志初始化
	logFile := utils.CreateLogFile()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.Println("日志初始化成功")
}

var db *gorm.DB

func init() {
	// 数据库初始化
	dsn := "user:passowrd@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	myDB, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err1 != nil {
		log.Fatal("数据库连接失败: ", err1)
	}
	log.Println("数据库连接成功")
	db = myDB
	err2 := db.AutoMigrate(&models.ToDoList{})
	if err2 != nil {
		log.Fatal("数据表迁移失败: ", err2)
	}
	log.Println("数据表迁移成功")
}

func main() {
	r := gin.Default()
	// 跨域
	r.Use(utils.Cors())
	// 配置路由
	routers.AddRouter(db, r)
	routers.FinishRouter(db, r)
	routers.DeleteRouter(db, r)
	routers.GetRouter(db, r)
	err := r.Run()
	if err != nil {
		log.Fatal("服务启动失败: ", err)
	}
}
