// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"path/filepath"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/Auth"
)

func RefactorPathValidator(path string, identification *Auth.Identification) bool {
	if !filepath.IsAbs(path) {
		return false
	}
	userDirectoryPath := FullPathExpressor("", identification)
	if !strings.HasPrefix(path, userDirectoryPath) {
		return false
	}
	return true
}
