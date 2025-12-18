// Copyright (c) ZStack.io, Inc.

package param

// InspectBareMetal2ChassisDetailParam InspectBareMetal2Chassis detail param
type InspectBareMetal2ChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// InspectBareMetal2ChassisParam InspectBareMetal2Chassis request param
type InspectBareMetal2ChassisParam struct {
	BaseParam
	Params InspectBareMetal2ChassisDetailParam `json:"params"`
}
