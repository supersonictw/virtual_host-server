// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package Http

type Identification struct {
	Identity string `json:"identity"`
	DisplayName string `json:"displayName"`
	PictureURL string  `json:"pictureUrl"`
	Email string `json:"email"`
}
