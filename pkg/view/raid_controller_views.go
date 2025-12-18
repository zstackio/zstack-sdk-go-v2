// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RaidControllerInventoryView RaidController
type RaidControllerInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"productName,omitempty"`
	rest string `json:"sasAddress,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int `json:"adapterNumber,omitempty"`
	rest []RaidPhysicalDriveInventoryView `json:"raidPhysicalDrives,omitempty"`
}

