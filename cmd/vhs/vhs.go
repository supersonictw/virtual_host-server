// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	vhsHttp "github.com/supersonictw/virtual_host-server/internal/http"
	"github.com/supersonictw/virtual_host-server/internal/user"
	"github.com/supersonictw/virtual_host-server/internal/user/fs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("FRONTEND_DOMAIN")}

	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "virtual_host-system",
			"copyright":   "(c)2021 SuperSonic(https://github.com/supersonictw)",
		})
	})

	router.GET("/authorize/:accessToken", func(c *gin.Context) {
		session, _ := vhsHttp.ReadAuthCookie(c)
		if session != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status": 403,
			})
			return
		}
		accessToken := c.Param("accessToken")
		err := vhsHttp.IssueAuthCookie(accessToken, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"reason": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})

	router.GET("/profile", func(c *gin.Context) {
		session := user.NewAccess(c)
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
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewRead(session, path)
		result := handler.Refactor().(*fs.ReadResponse)
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
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewMkdir(session, path)
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
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewWrite(session, path)
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
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewRemove(session, path)
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
