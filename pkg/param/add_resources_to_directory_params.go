// Copyright (c) ZStack.io, Inc.

package param

// AddResourcesToDirectoryDetailParam AddResourcesToDirectory detail param
type AddResourcesToDirectoryDetailParam struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// AddResourcesToDirectoryParam AddResourcesToDirectory request param
type AddResourcesToDirectoryParam struct {
	BaseParam
	Params AddResourcesToDirectoryDetailParam `json:"params"`
}
