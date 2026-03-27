// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2InstanceInventoryView BareMetal2Instance
type BareMetal2InstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	ChassisUuid string `json:"chassisUuid,omitempty"`
	LastChassisUuid string `json:"lastChassisUuid,omitempty"`
	GatewayUuid string `json:"gatewayUuid,omitempty"`
	LastGatewayUuid string `json:"lastGatewayUuid,omitempty"`
	ChassisOfferingUuid string `json:"chassisOfferingUuid,omitempty"`
	GatewayAllocatorStrategy string `json:"gatewayAllocatorStrategy,omitempty"`
	Status string `json:"status,omitempty"`
	ProvisionType string `json:"provisionType,omitempty"`
	AgentVersion string `json:"agentVersion,omitempty"`
	IsLatestAgent bool `json:"isLatestAgent,omitempty"`
	ProvisionNics []BareMetal2InstanceProvisionNicInventoryView `json:"provisionNics,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	LastHostUuid string `json:"lastHostUuid,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	RootVolumeUuid string `json:"rootVolumeUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Type string `json:"type,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	CpuSpeed int64 `json:"cpuSpeed,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
	State string `json:"state,omitempty"`
	VmNics []VmNicInventoryView `json:"vmNics,omitempty"`
	AllVolumes []VolumeInventoryView `json:"allVolumes,omitempty"`
	VmCdRoms []VmCdRomInventoryView `json:"vmCdRoms,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
}

// AttachProvisionNicToBondingEventView AttachProvisionNicToBondingEvent
type AttachProvisionNicToBondingEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// DetachProvisionNicFromBondingEventView DetachProvisionNicFromBondingEvent
type DetachProvisionNicFromBondingEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// QueryBareMetal2InstanceView QueryBareMetal2Instance
type QueryBareMetal2InstanceView struct {
	Inventories []BareMetal2InstanceInventoryView `json:"inventories,omitempty"`
}

// CreateBareMetal2InstanceEventView CreateBareMetal2InstanceEvent
type CreateBareMetal2InstanceEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// ChangeBareMetal2InstancePasswordEventView ChangeBareMetal2InstancePasswordEvent
type ChangeBareMetal2InstancePasswordEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// CreateBareMetal2InstanceFromVolumeBackupEventView CreateBareMetal2InstanceFromVolumeBackupEvent
type CreateBareMetal2InstanceFromVolumeBackupEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// ReconnectBareMetal2InstanceEventView ReconnectBareMetal2InstanceEvent
type ReconnectBareMetal2InstanceEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// StartBareMetal2InstanceEventView StartBareMetal2InstanceEvent
type StartBareMetal2InstanceEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// CreateBareMetal2InstanceFromVmBackupEventView CreateBareMetal2InstanceFromVmBackupEvent
type CreateBareMetal2InstanceFromVmBackupEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// UpdateBareMetal2InstanceEventView UpdateBareMetal2InstanceEvent
type UpdateBareMetal2InstanceEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

