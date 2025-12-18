// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmCdRomInventoryView VmCdRom
type VmCdRomInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest string `json:"isoUuid,omitempty"`
	rest string `json:"isoInstallPath,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

