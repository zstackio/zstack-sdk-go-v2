// Copyright (c) ZStack.io, Inc.

package param

// SetNicQosDetailParam SetNicQos详细参数
type SetNicQosDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int64 `json:"outboundBandwidth,omitempty"`
	rest int64 `json:"inboundBandwidth,omitempty"`
}

// SetNicQosParam SetNicQos请求参数
type SetNicQosParam struct {
	BaseParam
	Params SetNicQosDetailParam `json:"params"` // 详细参数
}

