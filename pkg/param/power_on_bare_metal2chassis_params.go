// Copyright (c) ZStack.io, Inc.

package param

// PowerOnBareMetal2ChassisDetailParam PowerOnBareMetal2Chassis detail param
type PowerOnBareMetal2ChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BootDev string `json:"bootDev,omitempty"`
}

// PowerOnBareMetal2ChassisParam PowerOnBareMetal2Chassis request param
type PowerOnBareMetal2ChassisParam struct {
	BaseParam
	Params PowerOnBareMetal2ChassisDetailParam `json:"params"`
}
