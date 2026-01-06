// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmInstanceDeviceAddressArchiveInventoryView VmInstanceDeviceAddressArchive
type VmInstanceDeviceAddressArchiveInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DeviceAddress string `json:"deviceAddress,omitempty"`
	AddressGroupUuid string `json:"addressGroupUuid,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	MetadataClass string `json:"metadataClass,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryVmInstanceDeviceAddressArchiveView QueryVmInstanceDeviceAddressArchive
type QueryVmInstanceDeviceAddressArchiveView struct {
	Inventories []VmInstanceDeviceAddressArchiveInventoryView `json:"inventories,omitempty"`
}

