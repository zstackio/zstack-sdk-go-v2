// Copyright (c) ZStack.io, Inc.

package param

// MoveResourcesToDirectoryDetailParam MoveResourcesToDirectory detail param
type MoveResourcesToDirectoryDetailParam struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// MoveResourcesToDirectoryParam MoveResourcesToDirectory request param
type MoveResourcesToDirectoryParam struct {
	BaseParam
	Params MoveResourcesToDirectoryDetailParam `json:"params"`
}
