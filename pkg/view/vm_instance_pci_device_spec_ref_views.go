// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmInstancePciDeviceSpecRefInventoryView VmInstancePciDeviceSpecRef
type VmInstancePciDeviceSpecRefInventoryView struct {
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"pciSpecUuid,omitempty"`
	rest int `json:"pciDeviceNumber,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

