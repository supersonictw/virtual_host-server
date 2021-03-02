// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"encoding/base64"
	"encoding/json"
	"github.com/supersonictw/virtual_host-server/internal/User/FileSystem/middleware"
	"github.com/supersonictw/virtual_host-server/internal/Http"
	"io/ioutil"
)

type ReadResponse struct {
	Status bool
	Data   string
}

type Read struct {
	session *Http.Session
	path    string
}

func NewRead(session *Http.Session, path string) Interface {
	instance := new(Read)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path)
	return instance
}

func (r *Read) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session) {
		return false
	}
	return true
}

func (r *Read) directoryHandler(response *ReadResponse) {
	directory, err := ioutil.ReadDir(r.path)
	if err != nil {
		panic(err)
	}
	json, err := json.Marshal(directory)
	if err != nil {
		panic(err)
	}
	response.Status = true
	response.Data = string(json)
}

func (r *Read) fileHandler(response *ReadResponse) {
	content, err := ioutil.ReadFile(r.path)
	if err != nil {
		panic(err)
	}
	response.Status = true
	response.Data = base64.StdEncoding.EncodeToString(content)
}

func (r *Read) Refactor() interface{} {
	response := new(ReadResponse)
	response.Status = false
	if !r.Validate() {
		return response
	}
	fileType := middleware.PathTypeDetector(r.path)
	switch fileType {
	case middleware.Directory:
		r.directoryHandler(response)
		break
	case middleware.File:
		r.fileHandler(response)
		break
	default:
		return fileType
	}
	return response
}
