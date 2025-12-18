// Copyright (c) ZStack.io, Inc.

package param

// DeleteNasFileSystemDetailParam DeleteNasFileSystem detail param
type DeleteNasFileSystemDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNasFileSystemParam DeleteNasFileSystem request param
type DeleteNasFileSystemParam struct {
	BaseParam
	Params DeleteNasFileSystemDetailParam `json:"params"`
}
