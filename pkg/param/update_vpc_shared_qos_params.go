// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcSharedQosDetailParam UpdateVpcSharedQos detail param
type UpdateVpcSharedQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcSharedQosParam UpdateVpcSharedQos request param
type UpdateVpcSharedQosParam struct {
	BaseParam
	Params UpdateVpcSharedQosDetailParam `json:"params"`
}
