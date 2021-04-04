// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supersonictw/virtual_host-server/internal/Auth"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func ReadAuthCookie(c *gin.Context) *Session {
	accessString, err := c.Cookie("vhs_access")
	if err != nil {
		return nil
	}
	if strings.Trim(accessString, " ") == "" {
		return nil
	}

	session := new(Session)
	return session
}

func IssueAuthCookie(accessToken string, c *gin.Context) {
	identification := Auth.NewAuthorization(accessToken).GetIdentification()
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, identification)
	token, err := tokenClaims.SignedString(os.Getenv("JWT_SALT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie("vhs_access", token, 3600, "/", "localhost", false, true)
}
