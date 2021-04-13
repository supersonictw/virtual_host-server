// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Session struct {
	Identification *Identification
	Context        *gin.Context
}

func (s *Session) Journalist(action string, target string) {
	logRootPath := os.Getenv("LOG_DIRECTORY_PATH")
	time := time.Now().Format("2006-01-02")
	logPath := fmt.Sprintf("%s/%s.log", logRootPath, time)
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	actionCode := fmt.Sprintf("[%s]", action)
	logger := log.New(file, actionCode, log.Ltime)
	logger.Printf(
		"%s (%s, %s, %s)\n",
		target,
		s.Identification.DisplayName,
		s.Identification.Email,
		s.Identification.Identity,
	)
}
