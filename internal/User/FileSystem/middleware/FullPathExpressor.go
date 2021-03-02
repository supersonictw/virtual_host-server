// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"fmt"
	"path/filepath"
	"github.com/supersonictw/virtual_host-server/internal/Http"
)

func FullPathExpressor(path string, identification *Http.Identification) string {
	identity := identification.Identity
	wordDirectory := fmt.Sprintf("%s/%s", identity, path)
	result, err := filepath.Abs(wordDirectory)
	if err != nil {
		panic(err)
	}
	return result
}