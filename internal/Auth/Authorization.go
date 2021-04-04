// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Auth

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	OpenID2 "golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Authorization struct {
	client   *oauth2.Service
	userInfo *oauth2.Userinfo
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func NewAuthorization(accessToken string) (*Authorization, error) {
	var err error
	instance := new(Authorization)
	ctx := context.Background()
	instance.client, err = oauth2.NewService(
		ctx,
		option.WithTokenSource(
			OpenID2.StaticTokenSource(
				&OpenID2.Token{
					AccessToken: accessToken,
				},
			),
		),
		option.WithScopes(
			oauth2.OpenIDScope,
			oauth2.UserinfoEmailScope,
			oauth2.UserinfoProfileScope,
		),
	)
	if err != nil {
		if os.Getenv("ENVIRONMENT") == "env" {
			panic(err)
		}
		return nil, err
	}
	instance.userInfo, err = instance.client.Userinfo.Get().Do()
	if err != nil {
		if os.Getenv("ENVIRONMENT") == "env" {
			panic(err)
		}
		return nil, err
	}
	return instance, nil
}

func (handler *Authorization) GetIdentification() *Identification {
	identification := new(Identification)
	identification.DisplayName = handler.userInfo.Name
	identification.Subject = handler.userInfo.Id
	identification.PictureURL = handler.userInfo.Picture
	identification.Email = handler.userInfo.Email
	return identification
}
