package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mgs": "haha",
		})
	})
	go func() {
		router.Run(":8080")
	}()

	// 如果想要接收到信号  kill -9 强杀命令
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 处理后续的逻辑
	fmt.Println("关闭server中。。。")
	fmt.Println("注销服务中。。。。")
}
