// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/virtual_host-server/internal/User"
	"github.com/supersonictw/virtual_host-server/internal/User/FileSystem"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "virtual_host-system",
			"copyright":   "(c)2021 SuperSonic(https://github.com/supersonictw)",
		})
	})

	router.GET("/user/:path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		handler := FileSystem.NewRead(session, path)
		result := handler.Refactor().(*FileSystem.ReadResponse)
		if result.Status {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
			})
		}
	})

	router.POST("/user/:path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		handler := FileSystem.NewMkdir(session, path)
		if handler.Refactor().(bool) {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
			})
		}
	})

	router.PUT("/user/:path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		handler := FileSystem.NewWrite(session, path)
		if handler.Refactor().(bool) {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
			})
		}
	})

	router.DELETE("/user/:path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		handler := FileSystem.NewRemove(session, path)
		if handler.Refactor().(bool) {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
			})
		}
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
