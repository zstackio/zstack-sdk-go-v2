// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkMtuDetailParam GetL3NetworkMtu详细参数
type GetL3NetworkMtuDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
}

// GetL3NetworkMtuParam GetL3NetworkMtu请求参数
type GetL3NetworkMtuParam struct {
	BaseParam
	Params GetL3NetworkMtuDetailParam `json:"params"` // 详细参数
}

