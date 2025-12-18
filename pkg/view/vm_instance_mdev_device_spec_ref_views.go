// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmInstanceMdevDeviceSpecRefInventoryView VmInstanceMdevDeviceSpecRef
type VmInstanceMdevDeviceSpecRefInventoryView struct {
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"mdevSpecUuid,omitempty"`
	rest int `json:"mdevDeviceNumber,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

