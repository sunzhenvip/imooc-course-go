package main

import (
	"go.uber.org/zap"
)

func main() {
	// logger, _ := zap.NewProduction()  // 使用生产环境用法
	logger, _ := zap.NewDevelopment() // 测试环境
	zap.NewDevelopmentConfig()
	defer logger.Sync() // 缓存刷新出来
	url := "https://imooc.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		zap.String("test_url", url),
		"url", url,
		"attempt", 3,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
