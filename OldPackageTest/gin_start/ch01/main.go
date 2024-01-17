package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID string `url:"id" json:"id" sd:"sd"`
}

func pong(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

	// c.JSON(http.StatusOK, map[string]string{
	// 	"message": "ssd",
	// })
}

func main() {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", pong)
		v1.POST("/post", getPost)
	}
	r.Run()
}

func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", 0)
	name := c.DefaultPostForm("")
}
