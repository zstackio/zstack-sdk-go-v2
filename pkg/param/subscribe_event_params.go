// Copyright (c) ZStack.io, Inc.

package param

// UpdateSubscribeEventDetailParam UpdateSubscribeEvent详细参数
type UpdateSubscribeEventDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"name,omitempty"`
}

// UpdateSubscribeEventParam UpdateSubscribeEvent请求参数
type UpdateSubscribeEventParam struct {
	BaseParam
	Params UpdateSubscribeEventDetailParam `json:"params"` // 详细参数
}

