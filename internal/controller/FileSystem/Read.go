// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"github.com/supersonictw/virtual_host-server/internal/controller/FileSystem/middleware"
	"github.com/supersonictw/virtual_host-server/internal/model"
	"io/ioutil"
)

type Read struct {
	session *model.Session
	path    string
}

func NewRead(session *model.Session, path string) Interface {
	instance := new(Read)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path)
	return instance
}

func (r *Read) validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session) {
		return false
	}
	return true
}

func (r *Read) refactor() interface{} {
	if !r.validate() {
		return false
	}
	content, err := ioutil.ReadFile(r.path)
	if err != nil {
		panic(err)
	}
	return content
}
