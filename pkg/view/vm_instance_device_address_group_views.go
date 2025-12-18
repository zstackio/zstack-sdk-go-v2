// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmInstanceDeviceAddressGroupInventoryView VmInstanceDeviceAddressGroup
type VmInstanceDeviceAddressGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VmInstanceDeviceAddressArchiveInventoryView `json:"addressList,omitempty"`
}

