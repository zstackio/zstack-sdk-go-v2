// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalHardwareInfoInventoryView BaremetalHardwareInfo
type BaremetalHardwareInfoInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"content,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

