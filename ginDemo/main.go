package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Const_log() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()

		since := time.Since(now)
		log.Printf("const time %d us\n", since/1000)

	}
}

func main() {
	r := gin.Default()
	r.Use(Const_log())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})

	r.Run()
}
