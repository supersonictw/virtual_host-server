// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

type Interface interface {
	Validate() bool
	Refactor() interface{}
}