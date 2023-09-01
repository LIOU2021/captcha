package main

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

// 这里可以换成redis
var store = base64Captcha.DefaultMemStore

func main() {
	id, b64s := GetCaptcha()
	fmt.Println(id, b64s)
	fmt.Println(verify("xdyQzMAh40ezxkxBK2Ca", "11286"))
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
