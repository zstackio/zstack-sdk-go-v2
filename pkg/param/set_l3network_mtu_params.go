// Copyright (c) ZStack.io, Inc.

package param

// SetL3NetworkMtuDetailParam SetL3NetworkMtu详细参数
type SetL3NetworkMtuDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest int `json:"mtu" validate:"required"` // 必填
}

// SetL3NetworkMtuParam SetL3NetworkMtu请求参数
type SetL3NetworkMtuParam struct {
	BaseParam
	Params SetL3NetworkMtuDetailParam `json:"params"` // 详细参数
}

