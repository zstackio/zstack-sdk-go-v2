// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateNasFileSystemParamDetail UpdateNasFileSystem detail param
type UpdateNasFileSystemParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateNasFileSystemParam UpdateNasFileSystem request param
type UpdateNasFileSystemParam struct {
	BaseParam
	UpdateNasFileSystem UpdateNasFileSystemParamDetail `json:"updateNasFileSystem"`
}
// DeleteNasFileSystemParamDetail DeleteNasFileSystem detail param
type DeleteNasFileSystemParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNasFileSystemParam DeleteNasFileSystem request param
type DeleteNasFileSystemParam struct {
	BaseParam
	DeleteNasFileSystem DeleteNasFileSystemParamDetail `json:"deleteNasFileSystem"`
}
