// Copyright (c) ZStack.io, Inc.

package param

// PowerOffBareMetal2ChassisDetailParam PowerOffBareMetal2Chassis detail param
type PowerOffBareMetal2ChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// PowerOffBareMetal2ChassisParam PowerOffBareMetal2Chassis request param
type PowerOffBareMetal2ChassisParam struct {
	BaseParam
	Params PowerOffBareMetal2ChassisDetailParam `json:"params"`
}
