// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2InstanceInventoryView BareMetal2Instance
type BareMetal2InstanceInventoryView struct {
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"lastChassisUuid,omitempty"`
	rest string `json:"gatewayUuid,omitempty"`
	rest string `json:"lastGatewayUuid,omitempty"`
	rest string `json:"chassisOfferingUuid,omitempty"`
	rest string `json:"gatewayAllocatorStrategy,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"provisionType,omitempty"`
	rest string `json:"agentVersion,omitempty"`
	rest bool `json:"isLatestAgent,omitempty"`
	rest BareMetal2InstanceProvisionNicInventoryView `json:"provisionNic,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"lastHostUuid,omitempty"`
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest string `json:"rootVolumeUuid,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"cpuSpeed,omitempty"`
	rest string `json:"allocatorStrategy,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"state,omitempty"`
	rest []VmNicInventoryView `json:"vmNics,omitempty"`
	rest []VolumeInventoryView `json:"allVolumes,omitempty"`
	rest []VmCdRomInventoryView `json:"vmCdRoms,omitempty"`
	rest string `json:"guestOsType,omitempty"`
}

