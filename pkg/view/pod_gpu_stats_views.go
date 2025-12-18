// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PodGpuStatsInventoryView PodGpuStats
type PodGpuStatsInventoryView struct {
	rest string `json:"podUuid,omitempty"`
	rest int `json:"gpuCount,omitempty"`
	rest int64 `json:"avgAllocatedMb,omitempty"`
	rest int64 `json:"totalGpuMemMb,omitempty"`
}

