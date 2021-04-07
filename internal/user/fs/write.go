// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"github.com/supersonictw/virtual_host-server/internal/User/fs/middleware"
	"github.com/supersonictw/virtual_host-server/internal/Http"
)

type Write struct {
	session *Http.Session
	path    string
}

func NewWrite(session *Http.Session, path string) Interface {
	instance := new(Write)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path, session.Identification)
	return instance
}

func (w *Write) Validate() bool {
	if !middleware.RefactorPathValidator(w.path, w.session.Identification) {
		return false
	}
	return true
}

func (w *Write) Refactor() interface{} {
	if !w.Validate() {
		return false
	}
	context := w.session.Context
	file, _ := context.FormFile("file")
	err := context.SaveUploadedFile(file, w.path)
	if err != nil {
		panic(err)
	}
	w.session.Journalist("Write", w.path)
	return true
}
