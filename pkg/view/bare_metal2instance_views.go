// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2InstanceInventoryView BareMetal2Instance
type BareMetal2InstanceInventoryView struct {
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
	ProvisionNic BareMetal2InstanceProvisionNicInventoryView `json:"provisionNic,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	State string `json:"state,omitempty"`
	VmNics []VmNicInventoryView `json:"vmNics,omitempty"`
	AllVolumes []VolumeInventoryView `json:"allVolumes,omitempty"`
	VmCdRoms []VmCdRomInventoryView `json:"vmCdRoms,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
}

