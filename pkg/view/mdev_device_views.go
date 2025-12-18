// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MdevDeviceInventoryView MdevDevice
type MdevDeviceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"parentUuid,omitempty"`
	rest string `json:"mttyUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"mdevSpecUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"chooser,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"vendor,omitempty"`
}

