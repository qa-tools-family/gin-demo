package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func loginEndpoint(c *gin.Context){
	fmt.Println("这是login方法")
}

func submitEndpoint(c *gin.Context){
	fmt.Println("这是submit方法")
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
	}

	v2 := r.Group("v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
	}

	r.Run()
}
