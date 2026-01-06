// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PrimaryStorageCapacityInventoryView PrimaryStorageCapacity
type PrimaryStorageCapacityInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity int64 `json:"systemUsedCapacity,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// GetPrimaryStorageCapacityView GetPrimaryStorageCapacity
type GetPrimaryStorageCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	Success bool `json:"success,omitempty"`
}

