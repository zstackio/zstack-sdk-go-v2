// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateVmInstancePciDeviceSpecRefParamDetail UpdateVmInstancePciDeviceSpecRef detail param
type UpdateVmInstancePciDeviceSpecRefParamDetail struct {
	PciDeviceNumber int `json:"pciDeviceNumber" validate:"required"`
}

// UpdateVmInstancePciDeviceSpecRefParam UpdateVmInstancePciDeviceSpecRef request param
type UpdateVmInstancePciDeviceSpecRefParam struct {
	BaseParam
	Params UpdateVmInstancePciDeviceSpecRefParamDetail `json:"updateVmInstancePciDeviceSpecRef"`
}
