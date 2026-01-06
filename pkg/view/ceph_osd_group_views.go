// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CephOsdGroupInventoryView CephOsdGroup
type CephOsdGroupInventoryView struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	Osds string `json:"osds,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// QueryCephOsdGroupView QueryCephOsdGroup
type QueryCephOsdGroupView struct {
	Inventories []CephOsdGroupInventoryView `json:"inventories,omitempty"`
}

