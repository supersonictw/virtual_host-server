// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"github.com/supersonictw/virtual_host-server/internal/Http"
	"path/filepath"
)


func RefactorPathValidator(path string, identification *Http.Identification) bool {
	if !filepath.IsAbs(path) {
		return false
	}
	// if !strings.HasPrefix(path, identification.Identity) {
	// 	return false
	// }
	return true
}