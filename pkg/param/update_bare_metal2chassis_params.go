// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2ChassisDetailParam UpdateBareMetal2Chassis detail param
type UpdateBareMetal2ChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisParam UpdateBareMetal2Chassis request param
type UpdateBareMetal2ChassisParam struct {
	BaseParam
	Params UpdateBareMetal2ChassisDetailParam `json:"params"`
}
