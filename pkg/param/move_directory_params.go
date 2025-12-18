// Copyright (c) ZStack.io, Inc.

package param

// MoveDirectoryDetailParam MoveDirectory详细参数
type MoveDirectoryDetailParam struct {
	rest string `json:"targetParentUuid" validate:"required"` // 必填
	rest string `json:"directoryUuid" validate:"required"` // 必填
}

// MoveDirectoryParam MoveDirectory请求参数
type MoveDirectoryParam struct {
	BaseParam
	Params MoveDirectoryDetailParam `json:"params"` // 详细参数
}

