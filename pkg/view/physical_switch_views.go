// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PhysicalSwitchInventoryView PhysicalSwitch
type PhysicalSwitchInventoryView struct {
	BaseInfoView
	BaseTimeView
	Ip                string                            `json:"ip,omitempty"`
	Mac               string                            `json:"mac,omitempty"`
	Mode              string                            `json:"mode,omitempty"`
	SoftwareVersion   string                            `json:"softwareVersion,omitempty"`
	SdnControllerUuid string                            `json:"sdnControllerUuid,omitempty"`
	Ports             []PhysicalSwitchPortInventoryView `json:"ports,omitempty"`
}

// QueryPhysicalSwitchView QueryPhysicalSwitch
type QueryPhysicalSwitchView struct {
	Inventories []PhysicalSwitchInventoryView `json:"inventories,omitempty"`
}
