// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

import (
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

func ReadAuthCookie(c *gin.Context) (*Session, error) {
	token, err := c.Cookie("vhs_access")
	if err != nil || strings.Trim(token, " ") == "" {
		return nil, jwt.ErrInvalidKey
	}

	secret := []byte(os.Getenv("JWT_SALT_SECRET"))
	session := new(Session)
	session.Context = c

	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Auth.Identification{},
		func(token *jwt.Token) (i interface{}, err error) {
			return secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*Auth.Identification)

	if !ok || !tokenClaims.Valid {
		return nil, jwt.ErrInvalidKey
	}

	session.Identification = claims

	return session, nil
}

func IssueAuthCookie(accessToken string, c *gin.Context) error {
	auth, err := Auth.NewAuthorization(accessToken)
	if err != nil {
		return err
	}
	identification := auth.GetIdentification()
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, identification)
	secret := []byte(os.Getenv("JWT_SALT_SECRET"))
	token, err := tokenClaims.SignedString(secret)
	if err != nil {
		return err
	}
	domain := os.Getenv("BACKEND_DOMAIN")
	secure := func() bool {
		status := os.Getenv("BACKEND_SSL")
		if status == "yes" {
			return true
		}
		return false
	}()
	c.SetCookie("vhs_access", token, 3600, "/", domain, secure, true)
	return nil
}
