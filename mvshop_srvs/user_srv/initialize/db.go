package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mvshop_srvs/user_srv/global"
	"os"
	"time"
)

func InitDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "root:root@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	c := global.ServerConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MysqlInfo.User, c.MysqlInfo.Password, c.MysqlInfo.Host, c.MysqlInfo.Port, c.MysqlInfo.Name)
	// 日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置不加s 表面
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("连接失败")
	}
	// db.AutoMigrate(&model.User{})
	// fmt.Println(genMd5("sdsdsdsdsdsd"))
}
