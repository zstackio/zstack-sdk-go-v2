// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostIpmiInventoryView HostIpmi
type HostIpmiInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ipmiAddress,omitempty"`
	rest string `json:"ipmiUsername,omitempty"`
	rest int `json:"ipmiPort,omitempty"`
	rest string `json:"ipmiPowerStatus,omitempty"`
}

