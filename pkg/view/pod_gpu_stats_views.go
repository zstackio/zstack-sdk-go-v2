// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PodGpuStatsInventoryView PodGpuStats
type PodGpuStatsInventoryView struct {
	BaseInfoView
	BaseTimeView
	PodUuid string `json:"podUuid,omitempty"`
	GpuCount int `json:"gpuCount,omitempty"`
	AvgAllocatedMb int64 `json:"avgAllocatedMb,omitempty"`
	TotalGpuMemMb int64 `json:"totalGpuMemMb,omitempty"`
}

