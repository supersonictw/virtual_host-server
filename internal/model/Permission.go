// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

type Permission struct {
	Access bool
	Create bool
	Upload bool
	Delete bool
	Admin bool
}