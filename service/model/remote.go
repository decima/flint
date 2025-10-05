package model

import "time"

type File struct {
	Name  string    `json:"name"`
	IsDir bool      `json:"is_dir"`
	Size  int64     `json:"size"`
	Mode  string    `json:"mode"`
	MTime time.Time `json:"mtime"`
}
