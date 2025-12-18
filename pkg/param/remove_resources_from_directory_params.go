// Copyright (c) ZStack.io, Inc.

package param

// RemoveResourcesFromDirectoryDetailParam RemoveResourcesFromDirectory详细参数
type RemoveResourcesFromDirectoryDetailParam struct {
	rest []string `json:"resourceUuids" validate:"required"` // 必填
	rest string `json:"directoryUuid" validate:"required"` // 必填
}

// RemoveResourcesFromDirectoryParam RemoveResourcesFromDirectory请求参数
type RemoveResourcesFromDirectoryParam struct {
	BaseParam
	Params RemoveResourcesFromDirectoryDetailParam `json:"params"` // 详细参数
}

