// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HbaDeviceInventoryView HbaDevice
type HbaDeviceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"hbaType,omitempty"`
	rest string `json:"createDate,omitempty"`
	rest string `json:"lastOpDate,omitempty"`
}

