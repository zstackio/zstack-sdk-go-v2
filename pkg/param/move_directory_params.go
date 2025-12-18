// Copyright (c) ZStack.io, Inc.

package param

// MoveDirectoryDetailParam MoveDirectory detail param
type MoveDirectoryDetailParam struct {
	TargetParentUuid string `json:"targetParentUuid" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// MoveDirectoryParam MoveDirectory request param
type MoveDirectoryParam struct {
	BaseParam
	Params MoveDirectoryDetailParam `json:"params"`
}
