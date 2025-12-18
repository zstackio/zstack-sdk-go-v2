// Copyright (c) ZStack.io, Inc.

package param

// GetLocalRaidPhysicalDriveSmartDetailParam GetLocalRaidPhysicalDriveSmart detail param
type GetLocalRaidPhysicalDriveSmartDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetLocalRaidPhysicalDriveSmartParam GetLocalRaidPhysicalDriveSmart request param
type GetLocalRaidPhysicalDriveSmartParam struct {
	BaseParam
	Params GetLocalRaidPhysicalDriveSmartDetailParam `json:"params"`
}
