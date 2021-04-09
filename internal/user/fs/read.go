// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"encoding/base64"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
	"io/fs"
	"io/ioutil"
)

type ReadResponse struct {
	Status bool        `json:"status"`
	Type   int         `json:"type"`
	Data   interface{} `json:"data"`
}

type File struct {
	Name         string `json:"name"`
	Type         int    `json:"type"`
	Size         int64  `json:"size"`
	Mode         string `json:"mode"`
	LastModified int64  `json:"lastModified"`
}

type Read struct {
	session *auth.Session
	path    string
}

func NewRead(session *auth.Session, path string) Interface {
	instance := new(Read)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	return instance
}

func (r *Read) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) {
		return false
	}
	return true
}

func getFilesInDirectory(files []fs.FileInfo) []*File {
	files_ := make([]*File, len(files))
	for i, f := range files {
		type_ := 0
		if f.IsDir() {
			type_ = 1
		}
		files_[i] = &File{
			Name:         f.Name(),
			Type:         type_,
			Size:         f.Size(),
			Mode:         f.Mode().String(),
			LastModified: f.ModTime().UnixNano(),
		}
	}
	return files_
}

func (r *Read) directoryHandler(response *ReadResponse) {
	directory, err := ioutil.ReadDir(r.path)
	if err != nil {
		panic(err)
	}
	response.Status = true
	response.Data = getFilesInDirectory(directory)
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
	r.session.Journalist("Read", r.path)
	return response
}
