// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ESXHostInventoryView ESXHost
type ESXHostInventoryView struct {
	rest string `json:"vCenterUuid,omitempty"`
	rest string `json:"morval,omitempty"`
	rest string `json:"esxiVersion,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest int64 `json:"totalCpuCapacity,omitempty"`
	rest int64 `json:"availableCpuCapacity,omitempty"`
	rest int `json:"cpuSockets,omitempty"`
	rest int64 `json:"totalMemoryCapacity,omitempty"`
	rest int64 `json:"availableMemoryCapacity,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest string `json:"ipmiAddress,omitempty"`
	rest string `json:"ipmiUsername,omitempty"`
	rest int `json:"ipmiPort,omitempty"`
	rest string `json:"ipmiPowerStatus,omitempty"`
	rest string `json:"cpuStatus,omitempty"`
	rest string `json:"memoryStatus,omitempty"`
	rest string `json:"diskStatus,omitempty"`
	rest string `json:"nicStatus,omitempty"`
	rest string `json:"gpuStatus,omitempty"`
	rest string `json:"powerSupplyStatus,omitempty"`
	rest string `json:"fanStatus,omitempty"`
	rest string `json:"raidStatus,omitempty"`
	rest string `json:"temperatureStatus,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

