// Copyright (c) ZStack.io, Inc.

package param

// AttachIsoToVmInstanceDetailParam AttachIsoToVmInstance detail param
type AttachIsoToVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	IsoUuid string `json:"isoUuid" validate:"required"`
}

// AttachIsoToVmInstanceParam AttachIsoToVmInstance request param
type AttachIsoToVmInstanceParam struct {
	BaseParam
	Params AttachIsoToVmInstanceDetailParam `json:"params"`
}
