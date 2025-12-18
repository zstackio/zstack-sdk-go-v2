// Copyright (c) ZStack.io, Inc.

package param

// DeleteSNSApplicationPlatformDetailParam DeleteSNSApplicationPlatform详细参数
type DeleteSNSApplicationPlatformDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSNSApplicationPlatformParam DeleteSNSApplicationPlatform请求参数
type DeleteSNSApplicationPlatformParam struct {
	BaseParam
	Params DeleteSNSApplicationPlatformDetailParam `json:"params"` // 详细参数
}

