// Copyright (c) ZStack.io, Inc.

package param

// AttachScsiLunToVmInstanceDetailParam AttachScsiLunToVmInstance detail param
type AttachScsiLunToVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	DisableMultiPathAttach bool `json:"disableMultiPathAttach,omitempty"`
}

// AttachScsiLunToVmInstanceParam AttachScsiLunToVmInstance request param
type AttachScsiLunToVmInstanceParam struct {
	BaseParam
	Params AttachScsiLunToVmInstanceDetailParam `json:"params"`
}
