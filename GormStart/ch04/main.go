package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int // 数据库中存储的字段 Company_id
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

// type User struct {
// 	UserId uint   `gorm:"primarykey"`
// 	Name   string `gorm:"column:user_name;type:varchar(50) not null;default:''"`
// }

// func (Product) TableName() string {
// 	return "test_user"
// }

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	// 日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshop",
		},
	})
	if err != nil {
		panic("数据库链接失败。。。")
	}
	// _ = db.AutoMigrate(&User{})

	var user User
	// db.Preload("Company").First(&user)
	db.Joins("Company").First(&user)
	fmt.Println(user.Company.Name)

}
