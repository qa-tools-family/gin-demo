package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	//r.Run()
	server := endless.NewServer(":8080", r)
	server.BeforeBegin = func(add string) {
		logrus.Printf("Actual pid is %d", syscall.Getpid())
		// save it somehow
	}
	server.ListenAndServe()
}
