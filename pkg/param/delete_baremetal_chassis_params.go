// Copyright (c) ZStack.io, Inc.

package param

// DeleteBaremetalChassisDetailParam DeleteBaremetalChassis detail param
type DeleteBaremetalChassisDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBaremetalChassisParam DeleteBaremetalChassis request param
type DeleteBaremetalChassisParam struct {
	BaseParam
	Params DeleteBaremetalChassisDetailParam `json:"params"`
}
