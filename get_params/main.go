package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

// 从 url 中获取参数
func getUrlData(c *gin.Context) {
	// 获取 url 参数
	name := c.Param("name")
	// c.String 支持自定义字符串返回
	c.String(http.StatusOK, "Hello %s", name)
}

// 从 params 中获取参数
func getParamData(c *gin.Context) {
	name := c.DefaultQuery("name", "world")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

// 从 form 表单中获取参数
func getFormData(c *gin.Context) {
	nick := c.DefaultPostForm("nick", "anonymous")
	message := c.PostForm("message")
	names := c.PostFormMap("names")
	logger.Info("names: ", names)
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func getJsonData(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2. 参数校验
	if len(product.Name) > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product %s must 20 characters", product.Name)})
		return
	}
	// 3. 逻辑处理
	product.CreatedAt = time.Now()
	log.Printf("Register product %s success", product.Name)

	// 4. 返回结果
	c.JSON(http.StatusOK, product)
}


func main()  {
	r := gin.Default()
	r.GET("/user/:name", getUrlData)
	r.GET("/user", getParamData)
	r.POST("/user", getFormData)
	r.POST("/data", getJsonData)

	r.Run()
}
