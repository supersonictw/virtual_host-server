// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

type Read struct {}

func NewRead() Interface {
	instance := new(Read)
	return instance
}

func (r *Read) validate() {}

func (r *Read) refactor(){}
