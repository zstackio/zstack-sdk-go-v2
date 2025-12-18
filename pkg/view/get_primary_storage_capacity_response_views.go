// Copyright (c) ZStack.io, Inc.

package view

// GetPrimaryStorageCapacityView GetPrimaryStorageCapacity
type GetPrimaryStorageCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	Success bool `json:"success,omitempty"`
}

