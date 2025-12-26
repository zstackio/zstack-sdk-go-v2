// Copyright (c) ZStack.io, Inc.

package view

// GetCpuMemoryCapacityView GetCpuMemoryCapacity
type GetCpuMemoryCapacityView struct {
	TotalCpu int64 `json:"totalCpu,omitempty"`
	AvailableCpu int64 `json:"availableCpu,omitempty"`
	TotalMemory int64 `json:"totalMemory,omitempty"`
	AvailableMemory int64 `json:"availableMemory,omitempty"`
	ManagedCpuNum int64 `json:"managedCpuNum,omitempty"`
	CapacityData []CpuMemoryCapacityDataView `json:"capacityData,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Success bool `json:"success,omitempty"`
}

