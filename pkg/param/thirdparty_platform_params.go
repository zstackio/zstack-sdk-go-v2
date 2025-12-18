// Copyright (c) ZStack.io, Inc.

package param

// UpdateThirdpartyPlatformDetailParam UpdateThirdpartyPlatform详细参数
type UpdateThirdpartyPlatformDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"template,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"stateEvent,omitempty"`
	rest int64 `json:"lastSyncDateMills,omitempty"`
}

// UpdateThirdpartyPlatformParam UpdateThirdpartyPlatform请求参数
type UpdateThirdpartyPlatformParam struct {
	BaseParam
	Params UpdateThirdpartyPlatformDetailParam `json:"params"` // 详细参数
}

