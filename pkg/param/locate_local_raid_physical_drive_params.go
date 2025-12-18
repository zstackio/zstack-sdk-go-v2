// Copyright (c) ZStack.io, Inc.

package param

// LocateLocalRaidPhysicalDriveDetailParam LocateLocalRaidPhysicalDrive detail param
type LocateLocalRaidPhysicalDriveDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Locate bool `json:"locate,omitempty"`
}

// LocateLocalRaidPhysicalDriveParam LocateLocalRaidPhysicalDrive request param
type LocateLocalRaidPhysicalDriveParam struct {
	BaseParam
	Params LocateLocalRaidPhysicalDriveDetailParam `json:"params"`
}
