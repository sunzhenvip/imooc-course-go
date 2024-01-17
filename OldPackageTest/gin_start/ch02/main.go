package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User       string `json:"user" binding:"required,min=3,max=10"`
	// Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	// RePassword string `json:"password" binding:"required,eqfield=Password"`
}

func main() {
	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginForm LoginForm
		if err := c.ShouldBind(&loginForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "登陆成功",
		})
	})
	router.Run(":8083")
}
