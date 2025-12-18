// Copyright (c) ZStack.io, Inc.

package param

// GetNicQosDetailParam GetNicQos detail param
type GetNicQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ForceSync bool `json:"forceSync,omitempty"`
}

// GetNicQosParam GetNicQos request param
type GetNicQosParam struct {
	BaseParam
	Params GetNicQosDetailParam `json:"params"`
}
