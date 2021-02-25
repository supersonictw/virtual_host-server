// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

import (
	"context"
	"github.com/supersonictw/virtual_host-server/internal/model"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Authorize struct {
	client *oauth2.Service
	userInfo *oauth2.Userinfo
}

func NewAuthorize(accessToken string) *Authorize {
	var err error
	instance := new(Authorize)
	ctx := context.Background()
	instance.client, err = oauth2.NewService(ctx, option.WithScopes(oauth2.OpenIDScope))
	if err != nil {
		panic(err)
	}
	instance.userInfo, err = instance.client.Userinfo.Get().Do()
	if err != nil {
		panic(err)
	}
	return instance
}

func (handler *Authorize) GetUser() *model.User {
	user := new(model.User)
	user.DisplayName = handler.userInfo.Name
	user.Identity = handler.userInfo.Id
	user.Picture = handler.userInfo.Picture
	user.Email = handler.userInfo.Email
	return user
}
