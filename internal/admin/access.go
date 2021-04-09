// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"github.com/supersonictw/virtual_host-server/internal/http"
)

func NewAccess(c *gin.Context) *http.Session {
	accessToken := http.ReadAccessToken(c)
	if accessToken == "" {
		return nil
	}
	authorization, err := auth.NewAuthorization(accessToken)
	if err != nil {
		return nil
	}
	return authorization.GetSession(c)
}
