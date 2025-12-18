// Copyright (c) ZStack.io, Inc.

package param

// AddResourcesToDirectoryDetailParam AddResourcesToDirectory详细参数
type AddResourcesToDirectoryDetailParam struct {
	rest []string `json:"resourceUuids" validate:"required"` // 必填
	rest string `json:"directoryUuid" validate:"required"` // 必填
}

// AddResourcesToDirectoryParam AddResourcesToDirectory请求参数
type AddResourcesToDirectoryParam struct {
	BaseParam
	Params AddResourcesToDirectoryDetailParam `json:"params"` // 详细参数
}

