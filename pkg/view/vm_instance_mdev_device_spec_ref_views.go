// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmInstanceMdevDeviceSpecRefInventoryView VmInstanceMdevDeviceSpecRef
type VmInstanceMdevDeviceSpecRefInventoryView struct {
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	MdevSpecUuid string `json:"mdevSpecUuid,omitempty"`
	MdevDeviceNumber int `json:"mdevDeviceNumber,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

