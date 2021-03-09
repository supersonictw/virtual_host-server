// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package User

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/virtual_host-server/internal/Http"
)

func NewAccess(c *gin.Context) *Http.Session {
	accessToken, err := c.Cookie("vhs_access")
	if err != nil {
		return nil
	}
	if strings.Trim(accessToken, " ") == "" {
		return nil
	}
	authentication := Http.NewAuthorization(accessToken)
	return authentication.GetSession(c)
}
