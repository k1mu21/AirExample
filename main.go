package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/example", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	// r.GET("/example2", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "world",
	// 	})
	// })
	r.Run()
}
