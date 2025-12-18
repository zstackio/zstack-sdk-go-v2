// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ChassisDiskInventoryView BareMetal2ChassisDisk
type BareMetal2ChassisDiskInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest int64 `json:"diskSize,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"wwn,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

