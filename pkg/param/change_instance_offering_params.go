// Copyright (c) ZStack.io, Inc.

package param

// ChangeInstanceOfferingDetailParam ChangeInstanceOffering detail param
type ChangeInstanceOfferingDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid" validate:"required"`
}

// ChangeInstanceOfferingParam ChangeInstanceOffering request param
type ChangeInstanceOfferingParam struct {
	BaseParam
	Params ChangeInstanceOfferingDetailParam `json:"params"`
}
