// Copyright (c) ZStack.io, Inc.

package param

// SetVipQosDetailParam SetVipQos详细参数
type SetVipQosDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"port,omitempty"`
	rest int64 `json:"outboundBandwidth,omitempty"`
	rest int64 `json:"inboundBandwidth,omitempty"`
}

// SetVipQosParam SetVipQos请求参数
type SetVipQosParam struct {
	BaseParam
	Params SetVipQosDetailParam `json:"params"` // 详细参数
}

