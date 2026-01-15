// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateBareMetal2ChassisPciDeviceParamDetail UpdateBareMetal2ChassisPciDevice detail param
type UpdateBareMetal2ChassisPciDeviceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ChassisUuid string `json:"chassisUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisPciDeviceParam UpdateBareMetal2ChassisPciDevice request param
type UpdateBareMetal2ChassisPciDeviceParam struct {
	BaseParam
	UpdateBareMetal2ChassisPciDevice UpdateBareMetal2ChassisPciDeviceParamDetail `json:"updateBareMetal2ChassisPciDevice"`
}
