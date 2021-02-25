// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

type Create struct {}

func NewCreate() Interface {
	instance := new(Create)
	return instance
}

func (c *Create) validate() {}

func (c *Create) refactor(){}
