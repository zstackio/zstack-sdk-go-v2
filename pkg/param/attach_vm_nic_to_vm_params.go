// Copyright (c) ZStack.io, Inc.

package param

// AttachVmNicToVmDetailParam AttachVmNicToVm detail param
type AttachVmNicToVmDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// AttachVmNicToVmParam AttachVmNicToVm request param
type AttachVmNicToVmParam struct {
	BaseParam
	Params AttachVmNicToVmDetailParam `json:"params"`
}
