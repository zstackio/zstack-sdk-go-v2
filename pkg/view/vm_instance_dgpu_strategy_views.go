// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmInstanceDGpuStrategyInventoryView VmInstanceDGpuStrategy
type VmInstanceDGpuStrategyInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	GpuSpecUuid string `json:"gpuSpecUuid,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ShmemSize int64 `json:"shmemSize,omitempty"`
	GpuDeviceUuid string `json:"gpuDeviceUuid,omitempty"`
	Chooser string `json:"chooser,omitempty"`
	AutoDetachOnStop bool `json:"autoDetachOnStop,omitempty"`
}

