// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package FileSystem

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/Http"
	"github.com/supersonictw/virtual_host-server/internal/User/FileSystem/middleware"
)

type ReadResponse struct {
	Status bool `json:"status"`
	Type int `json:"type"`
	Data   string `json:"data"`
}

type Read struct {
	session *Http.Session
	path    string
}

func NewRead(session *Http.Session, path string) Interface {
	instance := new(Read)
	instance.session = session
	instance.path = middleware.FullPathExpressor(path, session.Identification)
	fmt.Println(instance.path)
	return instance
}

func (r *Read) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) {
		return false
	}
	return true
}

func getFileNamesInDirectory(files []fs.FileInfo) []string {
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name()
	}
	return names
}

func (r *Read) directoryHandler(response *ReadResponse) {
	directory, err := ioutil.ReadDir(r.path)
	if err != nil {
		panic(err)
	}
	fileNames := getFileNamesInDirectory(directory)
	fmt.Println(fileNames)
	if err != nil {
		panic(err)
	}
	response.Status = true
	response.Data = strings.Join(fileNames, ",")
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
	response.Type = middleware.PathTypeDetector(r.path)
	if !r.Validate() {
		return response
	}
	switch response.Type {
	case middleware.Directory:
		r.directoryHandler(response)
		break
	case middleware.File:
		r.fileHandler(response)
		break
	default:
		return response
	}
	return response
}
