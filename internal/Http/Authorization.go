// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

import (
	"os"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Authorization struct {
	client   *oauth2.Service
	userInfo *oauth2.Userinfo
}

func NewAuthorization(accessToken string) *Authorization {
	var err error
	if  err = godotenv.Load(); err != nil {
		panic(err)
	}
	instance := new(Authorization)
	ctx := context.Background()
	instance.client, err = oauth2.NewService(
		ctx,
		option.WithCredentialsFile(
			os.Getenv("GOOGLE_CREDENTIALS_FILE_PATH"),
		),
		option.WithScopes(
			oauth2.OpenIDScope,
			oauth2.UserinfoEmailScope,
			oauth2.UserinfoProfileScope,
		),
	)
	if err != nil {
		panic(err)
	}
	instance.userInfo, err = instance.client.Userinfo.Get().Do()
	if err != nil {
		panic(err)
	}
	return instance
}

func (handler *Authorization) GetSession(c *gin.Context) *Session {
	session := new(Session)
	session.Identification = handler.GetIdentification()
	session.Context = c
	return session
}

func (handler *Authorization) GetIdentification() *Identification {
	identification := new(Identification)
	identification.DisplayName = handler.userInfo.Name
	identification.Identity = handler.userInfo.Id
	identification.Picture = handler.userInfo.Picture
	identification.Email = handler.userInfo.Email
	return identification
}
