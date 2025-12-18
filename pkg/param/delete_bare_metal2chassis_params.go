// Copyright (c) ZStack.io, Inc.

package param

// DeleteBareMetal2ChassisDetailParam DeleteBareMetal2Chassis detail param
type DeleteBareMetal2ChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2ChassisParam DeleteBareMetal2Chassis request param
type DeleteBareMetal2ChassisParam struct {
	BaseParam
	Params DeleteBareMetal2ChassisDetailParam `json:"params"`
}
