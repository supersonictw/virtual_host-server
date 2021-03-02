// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"github.com/supersonictw/virtual_host-server/internal/controller/FileSystem/middleware"
	"github.com/supersonictw/virtual_host-server/internal/model"
	"os"
)

type Remove struct {
	session *model.Session
	path    string
}

func NewRemove(session *model.Session, path string) Interface {
	instance := new(Remove)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path)
	return instance
}

func (r *Remove) validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session) {
		return false
	}
	return true
}

func (r *Remove) refactor() interface{} {
	if !r.validate() {
		return false
	}
	err := os.Remove(r.path)
	if err != nil {
		panic(err)
	}
	return true
}
