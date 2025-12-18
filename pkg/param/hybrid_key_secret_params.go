// Copyright (c) ZStack.io, Inc.

package param

// DeleteHybridKeySecretDetailParam DeleteHybridKeySecret详细参数
type DeleteHybridKeySecretDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteHybridKeySecretParam DeleteHybridKeySecret请求参数
type DeleteHybridKeySecretParam struct {
	BaseParam
	Params DeleteHybridKeySecretDetailParam `json:"params"` // 详细参数
}

