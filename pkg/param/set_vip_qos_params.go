// Copyright (c) ZStack.io, Inc.

package param

// SetVipQosDetailParam SetVipQos detail param
type SetVipQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Port int `json:"port,omitempty"`
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
}

// SetVipQosParam SetVipQos request param
type SetVipQosParam struct {
	BaseParam
	Params SetVipQosDetailParam `json:"params"`
}
