// Copyright (c) ZStack.io, Inc.

package view

// GetBackupStorageCapacityView GetBackupStorageCapacity
type GetBackupStorageCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Success bool `json:"success,omitempty"`
}

