// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"fmt"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"os"
	"path/filepath"

	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

type Rename struct {
	session *auth.Session
	path    string
	newPath string
}

func NewRename(session *auth.Session, path string) Interface {
	instance := new(Rename)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	newName := session.Context.PostForm("name")
	newPath := fmt.Sprintf("%s/%s", filepath.Dir(path), newName)
	instance.newPath = middleware.FullPathExpression(newPath, session.Identification)
	return instance
}

func (r *Rename) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) ||
		!middleware.RefactorPathValidator(r.newPath, r.session.Identification) {
		return false
	}
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (r *Rename) Refactor() interface{} {
	if !r.Validate() {
		return false
	}
	err := os.Rename(r.path, r.newPath)
	if err != nil {
		panic(err)
	}
	r.session.Journalist("Rename", r.path)
	return true
}
