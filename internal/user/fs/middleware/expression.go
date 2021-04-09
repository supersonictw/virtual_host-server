// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/supersonictw/virtual_host-server/internal/auth"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func FullPathExpression(path string, identification *auth.Identification) string {
	prefix := UserDirectoryPrefix(identification)
	storageRootDirectoryPath := os.Getenv("STORAGE_ROOT_DIRECTORY_PATH")
	wordDirectory := fmt.Sprintf("%s/%s/%s", storageRootDirectoryPath, prefix, path)
	result, err := filepath.Abs(wordDirectory)
	if err != nil {
		panic(err)
	}
	return result
}
