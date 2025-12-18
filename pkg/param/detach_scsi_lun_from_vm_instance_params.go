// Copyright (c) ZStack.io, Inc.

package param

// DetachScsiLunFromVmInstanceDetailParam DetachScsiLunFromVmInstance detail param
type DetachScsiLunFromVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// DetachScsiLunFromVmInstanceParam DetachScsiLunFromVmInstance request param
type DetachScsiLunFromVmInstanceParam struct {
	BaseParam
	Params DetachScsiLunFromVmInstanceDetailParam `json:"params"`
}
