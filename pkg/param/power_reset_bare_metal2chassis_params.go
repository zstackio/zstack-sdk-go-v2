// Copyright (c) ZStack.io, Inc.

package param

// PowerResetBareMetal2ChassisDetailParam PowerResetBareMetal2Chassis detail param
type PowerResetBareMetal2ChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BootDev string `json:"bootDev,omitempty"`
}

// PowerResetBareMetal2ChassisParam PowerResetBareMetal2Chassis request param
type PowerResetBareMetal2ChassisParam struct {
	BaseParam
	Params PowerResetBareMetal2ChassisDetailParam `json:"params"`
}
