// Copyright (c) ZStack.io, Inc.

package param

// AddHybridKeySecretDetailParam AddHybridKeySecret详细参数
type AddHybridKeySecretDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"key" validate:"required"` // 必填
	rest string `json:"secret" validate:"required"` // 必填
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest bool `json:"sync,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddHybridKeySecretParam AddHybridKeySecret请求参数
type AddHybridKeySecretParam struct {
	BaseParam
	Params AddHybridKeySecretDetailParam `json:"params"` // 详细参数
}

