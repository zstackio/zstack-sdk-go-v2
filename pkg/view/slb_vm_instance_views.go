// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SlbVmInstanceInventoryView SlbVmInstance
type SlbVmInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	SlbGroupUuid              string                                 `json:"slbGroupUuid,omitempty"`
	ConfigTasks               []SlbVmInstanceConfigTaskInventoryView `json:"configTasks,omitempty"`
	ConfigVersion             int64                                  `json:"configVersion,omitempty"`
	PublicNetworkUuid         string                                 `json:"publicNetworkUuid,omitempty"`
	VirtualRouterVips         []string                               `json:"virtualRouterVips,omitempty"`
	ApplianceVmType           string                                 `json:"applianceVmType,omitempty"`
	ManagementNetworkUuid     string                                 `json:"managementNetworkUuid,omitempty"`
	DefaultRouteL3NetworkUuid string                                 `json:"defaultRouteL3NetworkUuid,omitempty"`
	Status                    string                                 `json:"status,omitempty"`
	AgentPort                 int                                    `json:"agentPort,omitempty"`
	HaStatus                  string                                 `json:"haStatus,omitempty"`
	ZoneUuid                  string                                 `json:"zoneUuid,omitempty"`
	ClusterUuid               string                                 `json:"clusterUuid,omitempty"`
	ImageUuid                 string                                 `json:"imageUuid,omitempty"`
	HostUuid                  string                                 `json:"hostUuid,omitempty"`
	LastHostUuid              string                                 `json:"lastHostUuid,omitempty"`
	InstanceOfferingUuid      string                                 `json:"instanceOfferingUuid,omitempty"`
	RootVolumeUuid            string                                 `json:"rootVolumeUuid,omitempty"`
	Platform                  string                                 `json:"platform,omitempty"`
	Architecture              string                                 `json:"architecture,omitempty"`
	DefaultL3NetworkUuid      string                                 `json:"defaultL3NetworkUuid,omitempty"`
	Type                      string                                 `json:"type,omitempty"`
	HypervisorType            string                                 `json:"hypervisorType,omitempty"`
	MemorySize                int64                                  `json:"memorySize,omitempty"`
	ReservedMemorySize        int64                                  `json:"reservedMemorySize,omitempty"`
	CpuNum                    int                                    `json:"cpuNum,omitempty"`
	CpuSpeed                  int64                                  `json:"cpuSpeed,omitempty"`
	AllocatorStrategy         string                                 `json:"allocatorStrategy,omitempty"`
	State                     string                                 `json:"state,omitempty"`
	VmNics                    []VmNicInventoryView                   `json:"vmNics,omitempty"`
	AllVolumes                []VolumeInventoryView                  `json:"allVolumes,omitempty"`
	VmCdRoms                  []VmCdRomInventoryView                 `json:"vmCdRoms,omitempty"`
	GuestOsType               string                                 `json:"guestOsType,omitempty"`
}

// CreateSlbInstanceEventView CreateSlbInstanceEvent
type CreateSlbInstanceEventView struct {
	Inventory SlbVmInstanceInventoryView `json:"inventory,omitempty"`
}

// QuerySlbVmInstanceView QuerySlbVmInstance
type QuerySlbVmInstanceView struct {
	Inventories []SlbVmInstanceInventoryView `json:"inventories,omitempty"`
}
