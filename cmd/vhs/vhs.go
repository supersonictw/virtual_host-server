// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supersonictw/virtual_host-server/internal/User"
	"github.com/supersonictw/virtual_host-server/internal/User/FileSystem"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "virtual_host-system",
			"copyright":   "(c)2021 SuperSonic(https://github.com/supersonictw)",
		})
	})

	router.GET("/profile", func(c *gin.Context) {
		session := User.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   session.Identification,
		})
	})

	router.GET("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := FileSystem.NewRead(session, path)
		result := handler.Refactor().(*FileSystem.ReadResponse)
		if result.Status {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   result,
			})
		} else if result.Type == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 404,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
			})
		}
	})

	router.POST("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
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

	router.PUT("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
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

	router.DELETE("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := User.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
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

	exposePort := fmt.Sprintf(":%s", os.Getenv("EXPOSE_PORT"))
	if err := router.Run(exposePort); err != nil {
		panic(err)
	}
}
