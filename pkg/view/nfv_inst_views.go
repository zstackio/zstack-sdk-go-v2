// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstInventoryView NfvInst
type NfvInstInventoryView struct {
	BaseInfoView
	BaseTimeView
	ConfigVersion int `json:"configVersion,omitempty"`
	NfvInstGroupUuid string `json:"nfvInstGroupUuid,omitempty"`
	NetOsDistro string `json:"netOsDistro,omitempty"`
	BaseOsDistro string `json:"baseOsDistro,omitempty"`
	ClusterStatus string `json:"clusterStatus,omitempty"`
	StatusDetail string `json:"statusDetail,omitempty"`
	ApplianceVmType string `json:"applianceVmType,omitempty"`
	ManagementNetworkUuid string `json:"managementNetworkUuid,omitempty"`
	DefaultRouteL3NetworkUuid string `json:"defaultRouteL3NetworkUuid,omitempty"`
	Status string `json:"status,omitempty"`
	AgentPort int `json:"agentPort,omitempty"`
	HaStatus string `json:"haStatus,omitempty"`
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

// CreateNfvInstEventView CreateNfvInstEvent
type CreateNfvInstEventView struct {
	Inventory NfvInstInventoryView `json:"inventory,omitempty"`
}

// QueryNfvInstView QueryNfvInst
type QueryNfvInstView struct {
	Inventories []NfvInstInventoryView `json:"inventories,omitempty"`
}

// ReconnectNfvInstEventView ReconnectNfvInstEvent
type ReconnectNfvInstEventView struct {
	Inventory ApplianceVmInventoryView `json:"inventory,omitempty"`
}

