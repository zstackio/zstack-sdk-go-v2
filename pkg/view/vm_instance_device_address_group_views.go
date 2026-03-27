// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmInstanceDeviceAddressGroupInventoryView VmInstanceDeviceAddressGroup
type VmInstanceDeviceAddressGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceUuid string `json:"resourceUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	AddressList []VmInstanceDeviceAddressArchiveInventoryView `json:"addressList,omitempty"`
}

// QueryVmInstanceDeviceAddressGroupView QueryVmInstanceDeviceAddressGroup
type QueryVmInstanceDeviceAddressGroupView struct {
	Inventories []VmInstanceDeviceAddressGroupInventoryView `json:"inventories,omitempty"`
}

