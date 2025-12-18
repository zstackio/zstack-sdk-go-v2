// Copyright (c) ZStack.io, Inc.

package param

// CreateResourceStackFromAppDetailParam CreateResourceStackFromApp详细参数
type CreateResourceStackFromAppDetailParam struct {
	rest string `json:"appUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest bool `json:"rollback,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest bool `json:"withoutAppInfo,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateResourceStackFromAppParam CreateResourceStackFromApp请求参数
type CreateResourceStackFromAppParam struct {
	BaseParam
	Params CreateResourceStackFromAppDetailParam `json:"params"` // 详细参数
}

