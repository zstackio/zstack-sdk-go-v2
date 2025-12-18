// Copyright (c) ZStack.io, Inc.

package view

// GetAppBuildSystemCapacityView GetAppBuildSystemCapacity
type GetAppBuildSystemCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Success bool `json:"success,omitempty"`
}

