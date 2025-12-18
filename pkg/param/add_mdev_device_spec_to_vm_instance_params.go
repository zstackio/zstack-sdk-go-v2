// Copyright (c) ZStack.io, Inc.

package param

// AddMdevDeviceSpecToVmInstanceDetailParam AddMdevDeviceSpecToVmInstance detail param
type AddMdevDeviceSpecToVmInstanceDetailParam struct {
	MdevSpecUuid string `json:"mdevSpecUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	MdevDeviceNumber int `json:"mdevDeviceNumber,omitempty"`
}

// AddMdevDeviceSpecToVmInstanceParam AddMdevDeviceSpecToVmInstance request param
type AddMdevDeviceSpecToVmInstanceParam struct {
	BaseParam
	Params AddMdevDeviceSpecToVmInstanceDetailParam `json:"params"`
}
