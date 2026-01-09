// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PhysicalSwitchInventoryView PhysicalSwitch
type PhysicalSwitchInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Mode *string `json:"mode,omitempty"`
	SoftwareVersion *string `json:"softwareVersion,omitempty"`
	SdnControllerUuid *string `json:"sdnControllerUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	Ports []PhysicalSwitchPortInventoryView `json:"ports,omitempty"`
}

// QueryPhysicalSwitchView QueryPhysicalSwitch
type QueryPhysicalSwitchView struct {
	Inventories []PhysicalSwitchInventoryView `json:"inventories,omitempty"`
}

