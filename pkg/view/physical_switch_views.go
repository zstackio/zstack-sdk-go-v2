// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PhysicalSwitchInventoryView PhysicalSwitch
type PhysicalSwitchInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"mac,omitempty"`
	rest string `json:"mode,omitempty"`
	rest string `json:"softwareVersion,omitempty"`
	rest string `json:"sdnControllerUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []PhysicalSwitchPortInventoryView `json:"ports,omitempty"`
}

