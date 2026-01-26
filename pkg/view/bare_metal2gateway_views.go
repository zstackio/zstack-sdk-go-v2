// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2GatewayInventoryView BareMetal2Gateway
type BareMetal2GatewayInventoryView struct {
	BaseInfoView
	BaseTimeView
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	ProvisionNic BareMetal2GatewayProvisionNicInventoryView `json:"provisionNic,omitempty"`
	Username string `json:"username,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	OsDistribution string `json:"osDistribution,omitempty"`
	OsRelease string `json:"osRelease,omitempty"`
	OsVersion string `json:"osVersion,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	TotalCpuCapacity int64 `json:"totalCpuCapacity,omitempty"`
	AvailableCpuCapacity int64 `json:"availableCpuCapacity,omitempty"`
	CpuSockets int `json:"cpuSockets,omitempty"`
	TotalMemoryCapacity int64 `json:"totalMemoryCapacity,omitempty"`
	AvailableMemoryCapacity int64 `json:"availableMemoryCapacity,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiPowerStatus string `json:"ipmiPowerStatus,omitempty"`
	CpuStatus string `json:"cpuStatus,omitempty"`
	MemoryStatus string `json:"memoryStatus,omitempty"`
	DiskStatus string `json:"diskStatus,omitempty"`
	NicStatus string `json:"nicStatus,omitempty"`
	GpuStatus string `json:"gpuStatus,omitempty"`
	PowerSupplyStatus string `json:"powerSupplyStatus,omitempty"`
	FanStatus string `json:"fanStatus,omitempty"`
	RaidStatus string `json:"raidStatus,omitempty"`
	TemperatureStatus string `json:"temperatureStatus,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

// ReconnectBareMetal2GatewayEventView ReconnectBareMetal2GatewayEvent
type ReconnectBareMetal2GatewayEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// UpdateBareMetal2GatewayEventView UpdateBareMetal2GatewayEvent
type UpdateBareMetal2GatewayEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// DetachBareMetal2GatewayFromClusterEventView DetachBareMetal2GatewayFromClusterEvent
type DetachBareMetal2GatewayFromClusterEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// ChangeBareMetal2GatewayClusterEventView ChangeBareMetal2GatewayClusterEvent
type ChangeBareMetal2GatewayClusterEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// QueryBareMetal2GatewayView QueryBareMetal2Gateway
type QueryBareMetal2GatewayView struct {
	Inventories []BareMetal2GatewayInventoryView `json:"inventories,omitempty"`
}

// ChangeBareMetal2GatewayStateEventView ChangeBareMetal2GatewayStateEvent
type ChangeBareMetal2GatewayStateEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// DeleteBareMetal2GatewayEventView DeleteBareMetal2GatewayEvent
type DeleteBareMetal2GatewayEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachBareMetal2GatewayToClusterEventView AttachBareMetal2GatewayToClusterEvent
type AttachBareMetal2GatewayToClusterEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// AddHostEventView AddHostEvent
type AddHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
}

