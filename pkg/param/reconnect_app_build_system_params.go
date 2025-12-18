// Copyright (c) ZStack.io, Inc.

package param

// ReconnectAppBuildSystemDetailParam ReconnectAppBuildSystem详细参数
type ReconnectAppBuildSystemDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectAppBuildSystemParam ReconnectAppBuildSystem请求参数
type ReconnectAppBuildSystemParam struct {
	BaseParam
	Params ReconnectAppBuildSystemDetailParam `json:"params"` // 详细参数
}

