// Copyright (c) ZStack.io, Inc.

package param

// UpdateNasFileSystemDetailParam UpdateNasFileSystem detail param
type UpdateNasFileSystemDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateNasFileSystemParam UpdateNasFileSystem request param
type UpdateNasFileSystemParam struct {
	BaseParam
	Params UpdateNasFileSystemDetailParam `json:"params"`
}
