// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

import (
	"github.com/gin-gonic/gin"
)

type Session struct {
	Identification *Identification
	Context *gin.Context
}