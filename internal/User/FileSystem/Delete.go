// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"github.com/supersonictw/virtual_host-server/internal/User/FileSystem/middleware"
	"github.com/supersonictw/virtual_host-server/internal/Http"
	"os"
)

type Remove struct {
	session *Http.Session
	path    string
}

func NewRemove(session *Http.Session, path string) Interface {
	instance := new(Remove)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path, session.Identification)
	return instance
}

func (r *Remove) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) {
		return false
	}
	return true
}

func (r *Remove) Refactor() interface{} {
	if !r.Validate() {
		return false
	}
	err := os.Remove(r.path)
	if err != nil {
		panic(err)
	}
	return true
}
