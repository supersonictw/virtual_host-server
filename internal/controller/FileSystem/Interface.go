// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import "github.com/supersonictw/virtual_host-server/internal/model"

type Interface interface {
	session *model.Session
	path string
	fullPath string
	validate()
	refactor()
}