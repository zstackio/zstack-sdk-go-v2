// Copyright (c) ZStack.io, Inc.

package param

// DeleteDirectoryDetailParam DeleteDirectory详细参数
type DeleteDirectoryDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteDirectoryParam DeleteDirectory请求参数
type DeleteDirectoryParam struct {
	BaseParam
	Params DeleteDirectoryDetailParam `json:"params"` // 详细参数
}

