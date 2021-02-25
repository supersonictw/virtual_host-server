// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

type Delete struct {}

func NewDelete() Interface {
	instance := new(Delete)
	return instance
}

func (d *Delete) validate() {}

func (d *Delete) refactor(){}
