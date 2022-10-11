package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:name", func(context *gin.Context) {
		name := context.Param("name")
		msg := "name: " + name
		context.JSON(200, gin.H{
			"msg": msg,
		})
	})
	router.Run(":3000")
}
