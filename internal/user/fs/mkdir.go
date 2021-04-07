// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"os"

	"github.com/supersonictw/virtual_host-server/internal/http"
	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

type Mkdir struct {
	session *http.Session
	path    string
}

func NewMkdir(session *http.Session, path string) Interface {
	instance := new(Mkdir)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path, session.Identification)
	return instance
}

func (m *Mkdir) Validate() bool {
	if !middleware.RefactorPathValidator(m.path, m.session.Identification) {
		return false
	}
	if _, err := os.Stat(m.path); !os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *Mkdir) Refactor() interface{} {
	if !m.Validate() {
		return false
	}
	err := os.MkdirAll(m.path, 0755)
	if err != nil {
		panic(err)
	}
	m.session.Journalist("Mkdir", m.path)
	return true
}
