// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"github.com/supersonictw/virtual_host-server/internal/model"
	"strings"
	"path/filepath"
)


func RefactorPathValidator(path string, session *model.Session) bool {
	if !filepath.IsAbs(path) {
		return false
	}
	if !strings.HasPrefix(path, session.User.Identity) {
		return false
	}
	return true
}