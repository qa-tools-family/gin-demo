package main

import "github.com/gin-gonic/gin"

func hello(c *gin.Context)  {
	if para := c.DefaultQuery("name", ""); para == "" {
		panic("panic error")
	}
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	// Default 方法得到的实例会自动开启 Logger 和 Recovery 两个中间件
	r := gin.New()

	// 在单个 url 上启用 gin.Logger() Middleware
	r.GET("/hello", gin.Logger(), hello)

	// 在 url 组上启用 gin.Recover() Middleware
	v2 := r.Group("/v2")
	v2.Use(gin.Recovery())
	{
		v2.GET("/hello", gin.Logger(), hello)
	}

	r.Run()
}
