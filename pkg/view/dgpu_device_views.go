// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DGpuDeviceInventoryView DGpuDevice
type DGpuDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	ParentGpuUuid string `json:"parentGpuUuid,omitempty"`
	GpuSpecUuid string `json:"gpuSpecUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	AllocatedMemory int64 `json:"allocatedMemory,omitempty"`
	ShmemSize int64 `json:"shmemSize,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	SmPercentLimit int `json:"smPercentLimit,omitempty"`
	Status string `json:"status,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

// QueryDGpuDeviceView QueryDGpuDevice
type QueryDGpuDeviceView struct {
	Inventories []DGpuDeviceInventoryView `json:"inventories,omitempty"`
}

