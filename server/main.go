package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // ログやリカバリのミドルウェアが自動で設定される
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})
	r.Run() // デフォルトではポート8080で起動
}
