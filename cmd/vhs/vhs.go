// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"application": "virtual_host-system",
			"copyright": "(c)2021 SuperSonic(https://github.com/supersonictw)",
		})
	})

	router.GET("/user/:params", func(c *gin.Context) {})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}