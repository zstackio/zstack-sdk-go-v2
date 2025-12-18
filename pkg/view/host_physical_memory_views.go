// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostPhysicalMemoryInventoryView HostPhysicalMemory
type HostPhysicalMemoryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"manufacturer,omitempty"`
	rest string `json:"size,omitempty"`
	rest string `json:"speed,omitempty"`
	rest string `json:"clockSpeed,omitempty"`
	rest string `json:"locator,omitempty"`
	rest string `json:"serialNumber,omitempty"`
	rest string `json:"rank,omitempty"`
	rest string `json:"voltage,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

