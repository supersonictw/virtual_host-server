// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"github.com/supersonictw/virtual_host-server/internal/controller/FileSystem/middleware"
	"github.com/supersonictw/virtual_host-server/internal/model"
	"io/ioutil"
)

type Write struct {
	session *model.Session
	path    string
}

func NewWrite(session *model.Session, path string) Interface {
	instance := new(Write)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path)
	return instance
}

func (w *Write) validate() bool {
	if !middleware.RefactorPathValidator(w.path, w.session) {
		return false
	}
	return true
}

func (w *Write) refactor() interface{} {
	if !w.validate() {
		return false
	}
	err := ioutil.WriteFile(w.path, []byte{}, 0644)
	if err != nil {
		panic(err)
	}
	return true
}
