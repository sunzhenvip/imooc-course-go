package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"mvshop_srvs/user_srv/model"
	"os"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	io.WriteString(Md5, code)
	// return md5.Sum([]byte(code))
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
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
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置不加s 表面
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("连接失败")
	}
	db.AutoMigrate(&model.User{})
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("bobby%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}
	// fmt.Println(genMd5("sdsdsdsdsdsd"))

	// Using the default options
	// salt, encodedPwd := password.Encode("generic password", nil)
	// check := password.Verify("generic password", salt, encodedPwd, nil)
	// fmt.Println(check) // true

	// // Using custom options
	// options := &password.Options{16, 100, 32, md5.New}
	// salt, encodedPwd := password.Encode("generic password", options)
	// newPassword := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	// fmt.Println(newPassword)
	// passwordInfo := strings.Split(newPassword, "$")
	// fmt.Println(passwordInfo)
	// check := password.Verify("generic password", salt, encodedPwd, options)
	// fmt.Println(check) // true

}
