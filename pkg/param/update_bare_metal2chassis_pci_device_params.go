// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2ChassisPciDeviceDetailParam UpdateBareMetal2ChassisPciDevice detail param
type UpdateBareMetal2ChassisPciDeviceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ChassisUuid string `json:"chassisUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisPciDeviceParam UpdateBareMetal2ChassisPciDevice request param
type UpdateBareMetal2ChassisPciDeviceParam struct {
	BaseParam
	Params UpdateBareMetal2ChassisPciDeviceDetailParam `json:"params"`
}
