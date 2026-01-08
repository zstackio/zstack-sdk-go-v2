// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmInstanceDeviceAddressGroupInventoryView VmInstanceDeviceAddressGroup
type VmInstanceDeviceAddressGroupInventoryView struct {
	Uuid           string                                        `json:"uuid,omitempty"`
	ResourceUuid   string                                        `json:"resourceUuid,omitempty"`
	VmInstanceUuid string                                        `json:"vmInstanceUuid,omitempty"`
	CreateDate     time.Time                                     `json:"createDate,omitempty"`
	LastOpDate     time.Time                                     `json:"lastOpDate,omitempty"`
	AddressList    []VmInstanceDeviceAddressArchiveInventoryView `json:"addressList,omitempty"`
}

// QueryVmInstanceDeviceAddressGroupView QueryVmInstanceDeviceAddressGroup
type QueryVmInstanceDeviceAddressGroupView struct {
	Inventories []VmInstanceDeviceAddressGroupInventoryView `json:"inventories,omitempty"`
}
