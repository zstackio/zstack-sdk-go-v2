// Copyright (c) ZStack.io, Inc.

package param

// GetVipQosDetailParam GetVipQos detail param
type GetVipQosDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
}

// GetVipQosParam GetVipQos request param
type GetVipQosParam struct {
	BaseParam
	Params GetVipQosDetailParam `json:"params"`
}
