// Copyright (c) ZStack.io, Inc.

package param

// GetNicQosDetailParam GetNicQos详细参数
type GetNicQosDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"forceSync,omitempty"`
}

// GetNicQosParam GetNicQos请求参数
type GetNicQosParam struct {
	BaseParam
	Params GetNicQosDetailParam `json:"params"` // 详细参数
}

