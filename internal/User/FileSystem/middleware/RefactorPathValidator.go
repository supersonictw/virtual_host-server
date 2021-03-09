// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"strings"
	"github.com/supersonictw/virtual_host-server/internal/Http"
	"path/filepath"
)


func RefactorPathValidator(path string, identification *Http.Identification) bool {
	if !filepath.IsAbs(path) {
		return false
	}
	userDirectoryPath := FullPathExpressor("", identification)
	if !strings.HasPrefix(path, userDirectoryPath) {
	    return false
	}
	return true
}