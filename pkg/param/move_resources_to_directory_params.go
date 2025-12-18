// Copyright (c) ZStack.io, Inc.

package param

// MoveResourcesToDirectoryDetailParam MoveResourcesToDirectory详细参数
type MoveResourcesToDirectoryDetailParam struct {
	rest []string `json:"resourceUuids" validate:"required"` // 必填
	rest string `json:"directoryUuid" validate:"required"` // 必填
}

// MoveResourcesToDirectoryParam MoveResourcesToDirectory请求参数
type MoveResourcesToDirectoryParam struct {
	BaseParam
	Params MoveResourcesToDirectoryDetailParam `json:"params"` // 详细参数
}

