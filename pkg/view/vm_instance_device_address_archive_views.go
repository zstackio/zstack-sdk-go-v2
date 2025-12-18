// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmInstanceDeviceAddressArchiveInventoryView VmInstanceDeviceAddressArchive
type VmInstanceDeviceAddressArchiveInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"deviceAddress,omitempty"`
	rest string `json:"addressGroupUuid,omitempty"`
	rest string `json:"metadata,omitempty"`
	rest string `json:"metadataClass,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

