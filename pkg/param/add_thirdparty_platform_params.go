// Copyright (c) ZStack.io, Inc.

package param

// AddThirdpartyPlatformDetailParam AddThirdpartyPlatform详细参数
type AddThirdpartyPlatformDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"template" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddThirdpartyPlatformParam AddThirdpartyPlatform请求参数
type AddThirdpartyPlatformParam struct {
	BaseParam
	Params AddThirdpartyPlatformDetailParam `json:"params"` // 详细参数
}

