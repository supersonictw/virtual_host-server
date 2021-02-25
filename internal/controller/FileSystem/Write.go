// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

type Write struct {}

func NewWrite() Interface {
	instance := new(Write)
	return instance
}

func (w *Write) validate() {}

func (w *Write) refactor(){}
