package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

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
	})
	if err != nil {
		panic("sdsdsds")
	}
	_ = db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Code:  "D42",
	// 	Price: 100,
	// })

	// db.Delete(&Product{}, 1)
	// 查询
	var product Product

	// db.First(&Product{}, 2)
	//
	// db.Model(&product).Update("Price")

	db.Model(&product).Updates(Product{Price: 123123, Code: sql.NullString{"", true}})

	// 如果去更新product 只设置了 price  200
	// db.Model(&product).Delete()
	// db.First(&Product{}, "code = ?", "D42")
	// 设置全局的logger  这个 logger 在我们执行每个sql语句的时候会打印每一行sql
	// sql才是最重要的，本着这个原则的尽量的给大家看到每个api背后

}
