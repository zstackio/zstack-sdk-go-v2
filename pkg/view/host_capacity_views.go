// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostCapacityInventoryView HostCapacity
type HostCapacityInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int64 `json:"totalMemory,omitempty"`
	rest int64 `json:"totalCpu,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int `json:"cpuSockets,omitempty"`
	rest int `json:"cpuCoreNum,omitempty"`
	rest int64 `json:"availableMemory,omitempty"`
	rest int64 `json:"availableCpu,omitempty"`
	rest int64 `json:"totalPhysicalMemory,omitempty"`
	rest int64 `json:"availablePhysicalMemory,omitempty"`
}

