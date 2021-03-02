// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Session struct {
	Identification *Identification
	Context        *gin.Context
}

func (s *Session) Journalist(action string, target string) {
	log.Printf("[%s] %s (%s, %s)\n", action, target, s.Identification.DisplayName, s.Identification.Identity)
}
