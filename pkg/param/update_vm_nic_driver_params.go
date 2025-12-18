// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmNicDriverDetailParam UpdateVmNicDriver detail param
type UpdateVmNicDriverDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	DriverType string `json:"driverType" validate:"required"`
}

// UpdateVmNicDriverParam UpdateVmNicDriver request param
type UpdateVmNicDriverParam struct {
	BaseParam
	Params UpdateVmNicDriverDetailParam `json:"params"`
}
