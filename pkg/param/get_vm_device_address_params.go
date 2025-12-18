// Copyright (c) ZStack.io, Inc.

package param

// GetVmDeviceAddressDetailParam GetVmDeviceAddress detail param
type GetVmDeviceAddressDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceTypes []string `json:"resourceTypes" validate:"required"`
}

// GetVmDeviceAddressParam GetVmDeviceAddress request param
type GetVmDeviceAddressParam struct {
	BaseParam
	Params GetVmDeviceAddressDetailParam `json:"params"`
}
