// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DGpuProfileInventoryView DGpuProfile
type DGpuProfileInventoryView struct {
	BaseInfoView
	BaseTimeView
	GpuSpecUuid string `json:"gpuSpecUuid,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ShmemSize int64 `json:"shmemSize,omitempty"`
}

// QueryDGpuProfileView QueryDGpuProfile
type QueryDGpuProfileView struct {
	Inventories []DGpuProfileInventoryView `json:"inventories,omitempty"`
}

// SetDGpuProfileEventView SetDGpuProfileEvent
type SetDGpuProfileEventView struct {
	Inventories []DGpuProfileInventoryView `json:"inventories,omitempty"`
}

