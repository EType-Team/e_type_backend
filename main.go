package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})
	router.Run() // 0.0.0.0:8080 でサーバーをたてる
}
