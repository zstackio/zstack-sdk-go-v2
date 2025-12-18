// Copyright (c) ZStack.io, Inc.

package param

// PowerOnBaremetalChassisDetailParam PowerOnBaremetalChassis detail param
type PowerOnBaremetalChassisDetailParam struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
}

// PowerOnBaremetalChassisParam PowerOnBaremetalChassis request param
type PowerOnBaremetalChassisParam struct {
	BaseParam
	Params PowerOnBaremetalChassisDetailParam `json:"params"`
}
