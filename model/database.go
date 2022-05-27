package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Mydb *gorm.DB //方便外面能使用到数据库
func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/webgo?charset=utf8mb4&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置一个日志
	})
	if err != nil {
		log.Panic("failed to connect database")
	}
	Mydb = d
	fmt.Println("数据库连接成功")
	return Mydb
}
