// Copyright (c) ZStack.io, Inc.

package param

// DeleteDirectoryDetailParam DeleteDirectory detail param
type DeleteDirectoryDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteDirectoryParam DeleteDirectory request param
type DeleteDirectoryParam struct {
	BaseParam
	Params DeleteDirectoryDetailParam `json:"params"`
}
