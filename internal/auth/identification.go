// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package auth

import "gopkg.in/dgrijalva/jwt-go.v3"

type Identification struct {
	jwt.StandardClaims
	DisplayName string `json:"name"`
	PictureURL  string `json:"picture"`
	Email       string `json:"email"`
}
