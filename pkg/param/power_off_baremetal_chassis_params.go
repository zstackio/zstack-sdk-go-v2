// Copyright (c) ZStack.io, Inc.

package param

// PowerOffBaremetalChassisDetailParam PowerOffBaremetalChassis detail param
type PowerOffBaremetalChassisDetailParam struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
}

// PowerOffBaremetalChassisParam PowerOffBaremetalChassis request param
type PowerOffBaremetalChassisParam struct {
	BaseParam
	Params PowerOffBaremetalChassisDetailParam `json:"params"`
}
