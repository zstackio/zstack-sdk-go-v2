// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RaidPhysicalDriveInventoryView RaidPhysicalDrive
type RaidPhysicalDriveInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"raidLevel,omitempty"`
	rest string `json:"raidControllerUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"enclosureDeviceId,omitempty"`
	rest int `json:"slotNumber,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest int `json:"diskGroup,omitempty"`
	rest string `json:"wwn,omitempty"`
	rest string `json:"serialNumber,omitempty"`
	rest string `json:"deviceModel,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"driveState,omitempty"`
	rest string `json:"locateStatus,omitempty"`
	rest string `json:"driveType,omitempty"`
	rest string `json:"mediaType,omitempty"`
	rest int `json:"rotationRate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

