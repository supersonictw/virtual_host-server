// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"github.com/joho/godotenv"
	"github.com/supersonictw/virtual_host-server/internal/Http"
	"strings"
	"sort"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func UserDirectoryPrefix(identification *Http.Identification) string {
	if os.Getenv("STORAGE_USER_DIRECTORY_NAME_METHOD") == "email" {
		splited := strings.Split(identification.Email, "@")
    	sort.Sort(sort.Reverse(sort.StringSlice(splited)))
		return strings.Join(splited, "/")
	}
	return identification.Identity
}
