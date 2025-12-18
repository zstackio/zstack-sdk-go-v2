// Copyright (c) ZStack.io, Inc.

package param

// UpdateDirectoryDetailParam UpdateDirectory detail param
type UpdateDirectoryDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// UpdateDirectoryParam UpdateDirectory request param
type UpdateDirectoryParam struct {
	BaseParam
	Params UpdateDirectoryDetailParam `json:"params"`
}
