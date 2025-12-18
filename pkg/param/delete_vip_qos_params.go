// Copyright (c) ZStack.io, Inc.

package param

// DeleteVipQosDetailParam DeleteVipQos detail param
type DeleteVipQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Port int `json:"port,omitempty"`
}

// DeleteVipQosParam DeleteVipQos request param
type DeleteVipQosParam struct {
	BaseParam
	Params DeleteVipQosDetailParam `json:"params"`
}
