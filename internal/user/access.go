// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/virtual_host-server/internal/Http"
)

func NewAccess(c *gin.Context) *Http.Session {
	session, err := Http.ReadAuthCookie(c)
	if err != nil {
		return nil
	}
	return session
}
