// Copyright (c) ZStack.io, Inc.

package param

// DetachHybridKeyDetailParam DetachHybridKey详细参数
type DetachHybridKeyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DetachHybridKeyParam DetachHybridKey请求参数
type DetachHybridKeyParam struct {
	BaseParam
	Params DetachHybridKeyDetailParam `json:"params"` // 详细参数
}

