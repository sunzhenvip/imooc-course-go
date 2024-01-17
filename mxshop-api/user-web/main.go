package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
	myvalidator "mxshop-api/user-web/validator"
)

func main() {
	// router := gin.Default()
	// ApiGroup := router.Group("/v1") // 加入版本号
	// router2.InitUserRouter(ApiGroup)
	// router.GET("/ping")
	// port := 8021
	// logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()
	// zap.ReplaceGlobals(logger)
	// 1、初始化logger
	initialize.InitLogger()

	initialize.InitConfig()
	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// srv 初始化服务连接
	initialize.InitSrvConn()
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		// 翻译自定义为中文描述
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法都手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	// defer logger.Sync()
	// 2、初始化routers
	Router := initialize.Routers()
	/**
	1、S 可以获取一个全局都Sugar 可以让我们自己设置一个全局都logger
	2、日志分级别 debug info ,warn ,error fetal
	3、S函数和L函数很有用 提供了一个全局都安全都访问logger都途径
	*/
	// zap.S().Infof("启动服务器,端口:%d", port)
	zap.S().Debugf("启动服务器,端口:%d", global.ServerConfig.Port)

	// 启动发现报错
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error)
	}
}
