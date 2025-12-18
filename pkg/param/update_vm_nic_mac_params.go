// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmNicMacDetailParam UpdateVmNicMac detail param
type UpdateVmNicMacDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	Mac string `json:"mac" validate:"required"`
}

// UpdateVmNicMacParam UpdateVmNicMac request param
type UpdateVmNicMacParam struct {
	BaseParam
	Params UpdateVmNicMacDetailParam `json:"params"`
}
