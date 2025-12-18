// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceFirstBootDeviceDetailParam GetVmInstanceFirstBootDevice detail param
type GetVmInstanceFirstBootDeviceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmInstanceFirstBootDeviceParam GetVmInstanceFirstBootDevice request param
type GetVmInstanceFirstBootDeviceParam struct {
	BaseParam
	Params GetVmInstanceFirstBootDeviceDetailParam `json:"params"`
}
