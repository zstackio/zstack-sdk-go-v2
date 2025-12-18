// Copyright (c) ZStack.io, Inc.

package param

// InspectBaremetalChassisDetailParam InspectBaremetalChassis detail param
type InspectBaremetalChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// InspectBaremetalChassisParam InspectBaremetalChassis request param
type InspectBaremetalChassisParam struct {
	BaseParam
	Params InspectBaremetalChassisDetailParam `json:"params"`
}
