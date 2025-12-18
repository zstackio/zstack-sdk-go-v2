// Copyright (c) ZStack.io, Inc.

package param

// DetachIsoFromVmInstanceDetailParam DetachIsoFromVmInstance detail param
type DetachIsoFromVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	IsoUuid string `json:"isoUuid,omitempty"`
}

// DetachIsoFromVmInstanceParam DetachIsoFromVmInstance request param
type DetachIsoFromVmInstanceParam struct {
	BaseParam
	Params DetachIsoFromVmInstanceDetailParam `json:"params"`
}
