// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"os"

	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

type Remove struct {
	session *auth.Session
	path    string
}

func NewRemove(session *auth.Session, path string) Interface {
	instance := new(Remove)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	return instance
}

func (r *Remove) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) {
		return false
	}
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (r *Remove) Refactor() interface{} {
	if !r.Validate() {
		return false
	}
	err := os.RemoveAll(r.path)
	if err != nil {
		panic(err)
	}
	r.session.Journalist("Remove", r.path)
	return true
}
