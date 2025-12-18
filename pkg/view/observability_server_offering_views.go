// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ObservabilityServerOfferingInventoryView ObservabilityServerOffering
type ObservabilityServerOfferingInventoryView struct {
	ManagementNetworkUuid string `json:"managementNetworkUuid,omitempty"`
	PublicNetworkUuid string `json:"publicNetworkUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	CpuSpeed int `json:"cpuSpeed,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	Type string `json:"type,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
	SortKey int `json:"sortKey,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	State string `json:"state,omitempty"`
}

