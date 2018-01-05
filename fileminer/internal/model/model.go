package model

import "time"

//FileInfo contains info, given by file miner
type FileInfo struct {
	ID      int       `json:"-" xml:"-"`
	Name    string    `json:"name" xml:"name"`
	Mode    string    `json:"mode" xml:"mode"`
	ModTime time.Time `json:"mod_time" xml:"mod_time"`
	Size    int64     `json:"size" xml:"size"`
}
