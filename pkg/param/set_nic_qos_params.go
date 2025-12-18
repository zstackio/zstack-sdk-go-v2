// Copyright (c) ZStack.io, Inc.

package param

// SetNicQosDetailParam SetNicQos detail param
type SetNicQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
}

// SetNicQosParam SetNicQos request param
type SetNicQosParam struct {
	BaseParam
	Params SetNicQosDetailParam `json:"params"`
}
