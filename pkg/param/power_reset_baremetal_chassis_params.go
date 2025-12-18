// Copyright (c) ZStack.io, Inc.

package param

// PowerResetBaremetalChassisDetailParam PowerResetBaremetalChassis detail param
type PowerResetBaremetalChassisDetailParam struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
}

// PowerResetBaremetalChassisParam PowerResetBaremetalChassis request param
type PowerResetBaremetalChassisParam struct {
	BaseParam
	Params PowerResetBaremetalChassisDetailParam `json:"params"`
}
