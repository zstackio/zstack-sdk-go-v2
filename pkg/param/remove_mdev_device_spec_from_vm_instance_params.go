// Copyright (c) ZStack.io, Inc.

package param

// RemoveMdevDeviceSpecFromVmInstanceDetailParam RemoveMdevDeviceSpecFromVmInstance detail param
type RemoveMdevDeviceSpecFromVmInstanceDetailParam struct {
	MdevSpecUuid string `json:"mdevSpecUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// RemoveMdevDeviceSpecFromVmInstanceParam RemoveMdevDeviceSpecFromVmInstance request param
type RemoveMdevDeviceSpecFromVmInstanceParam struct {
	BaseParam
	Params RemoveMdevDeviceSpecFromVmInstanceDetailParam `json:"params"`
}
