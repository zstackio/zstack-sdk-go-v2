// Copyright (c) ZStack.io, Inc.

package param

// DeleteNicQosDetailParam DeleteNicQos detail param
type DeleteNicQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Direction string `json:"direction" validate:"required"`
}

// DeleteNicQosParam DeleteNicQos request param
type DeleteNicQosParam struct {
	BaseParam
	Params DeleteNicQosDetailParam `json:"params"`
}
