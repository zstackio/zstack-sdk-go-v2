// Copyright (c) ZStack.io, Inc.

package param

// EjectZBoxDetailParam EjectZBox详细参数
type EjectZBoxDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// EjectZBoxParam EjectZBox请求参数
type EjectZBoxParam struct {
	BaseParam
	Params EjectZBoxDetailParam `json:"params"` // 详细参数
}

