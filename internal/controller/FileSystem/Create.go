// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"github.com/supersonictw/virtual_host-server/internal/controller/FileSystem/middleware"
	"github.com/supersonictw/virtual_host-server/internal/model"
	"os"
)

type Create struct {
	session *model.Session
	path    string
}

func NewCreate(session *model.Session, path string) Interface {
	instance := new(Create)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path)
	return instance
}

func (c *Create) validate() bool {
	if !middleware.RefactorPathValidator(c.path, c.session) {
		return false
	}
	return true
}

func (c *Create) refactor() interface{} {
	if !c.validate() {
		return false
	}
	err := os.MkdirAll(c.path, 0755)
	if err != nil {
		panic(err)
	}
	return true
}
