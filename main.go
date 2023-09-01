package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// 这里可以换成redis
var store = base64Captcha.DefaultMemStore

func main() {
	r := gin.Default()
	r.Static("demo", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("getCode", func(c *gin.Context) {
		id, b64s := GetCaptcha()
		c.JSON(200, gin.H{
			"id":   id,
			"data": b64s,
		})
	})

	r.GET("verifyCode", func(c *gin.Context) {
		c.String(200, "%v", verify(c.Query("codeId"), c.Query("code")))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 获取验证马
func GetCaptcha() (string, string) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	c := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := c.Generate()

	if err != nil {
		fmt.Println("register error")
		return "", ""
	}

	return id, b64s
}

func verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}

	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}
