// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/joho/godotenv"
	"github.com/supersonictw/virtual_host-server/internal/Http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func FullPathExpressor(path string, identification *Http.Identification) string {
	identity := identification.Identity
	storageRootDirectoryPath := os.Getenv("STORAGE_ROOT_DIRECTORY_PATH")
	wordDirectory := fmt.Sprintf("%s/%s/%s", storageRootDirectoryPath, identity, path)
	result, err := filepath.Abs(wordDirectory)
	if err != nil {
		panic(err)
	}
	return result
}