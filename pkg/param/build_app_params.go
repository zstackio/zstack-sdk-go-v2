// Copyright (c) ZStack.io, Inc.

package param

// DeleteBuildAppDetailParam DeleteBuildApp详细参数
type DeleteBuildAppDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteBuildAppParam DeleteBuildApp请求参数
type DeleteBuildAppParam struct {
	BaseParam
	Params DeleteBuildAppDetailParam `json:"params"` // 详细参数
}

