// Copyright (c) ZStack.io, Inc.

package param

// RemoveResourcesFromDirectoryDetailParam RemoveResourcesFromDirectory detail param
type RemoveResourcesFromDirectoryDetailParam struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// RemoveResourcesFromDirectoryParam RemoveResourcesFromDirectory request param
type RemoveResourcesFromDirectoryParam struct {
	BaseParam
	Params RemoveResourcesFromDirectoryDetailParam `json:"params"`
}
