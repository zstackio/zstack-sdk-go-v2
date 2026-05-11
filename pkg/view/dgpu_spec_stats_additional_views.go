// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DGpuSpecStatsInventoryView DGpuSpecStats
type DGpuSpecStatsInventoryView struct {
	BaseInfoView
	BaseTimeView
	GpuSpecUuid string `json:"gpuSpecUuid,omitempty"`
	GpuSpecName string `json:"gpuSpecName,omitempty"`
	GpuType string `json:"gpuType,omitempty"`
	GpuCount int64 `json:"gpuCount,omitempty"`
	DgpuCount int64 `json:"dgpuCount,omitempty"`
	TotalMemory int64 `json:"totalMemory,omitempty"`
	AllocatedMemory int64 `json:"allocatedMemory,omitempty"`
	AvailableMemory int64 `json:"availableMemory,omitempty"`
	AllocationRate float64 `json:"allocationRate,omitempty"`
}

