// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RaidControllerInventoryView RaidController
type RaidControllerInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	ProductName string `json:"productName,omitempty"`
	SasAddress string `json:"sasAddress,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AdapterNumber int `json:"adapterNumber,omitempty"`
	RaidPhysicalDrives []RaidPhysicalDriveInventoryView `json:"raidPhysicalDrives,omitempty"`
}

