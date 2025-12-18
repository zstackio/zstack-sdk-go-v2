// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostHwMonitorStatusInventoryView HostHwMonitorStatus
type HostHwMonitorStatusInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"cpuStatus,omitempty"`
	rest string `json:"memoryStatus,omitempty"`
	rest string `json:"diskStatus,omitempty"`
	rest string `json:"fanStatus,omitempty"`
	rest string `json:"powerSupplyStatus,omitempty"`
	rest string `json:"raidStatus,omitempty"`
	rest string `json:"networkStatus,omitempty"`
	rest string `json:"gpuStatus,omitempty"`
	rest string `json:"temperatureStatus,omitempty"`
}

