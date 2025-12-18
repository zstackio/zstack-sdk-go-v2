// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcSharedQosDetailParam DeleteVpcSharedQos detail param
type DeleteVpcSharedQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcSharedQosParam DeleteVpcSharedQos request param
type DeleteVpcSharedQosParam struct {
	BaseParam
	Params DeleteVpcSharedQosDetailParam `json:"params"`
}
