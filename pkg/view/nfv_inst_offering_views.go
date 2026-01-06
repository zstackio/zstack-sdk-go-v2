// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstOfferingInventoryView NfvInstOffering
type NfvInstOfferingInventoryView struct {
	ManagementNetworkUuid string `json:"managementNetworkUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
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
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	State string `json:"state,omitempty"`
}

// QueryNfvInstOfferingView QueryNfvInstOffering
type QueryNfvInstOfferingView struct {
	Inventories []NfvInstOfferingInventoryView `json:"inventories,omitempty"`
}

// CreateInstanceOfferingEventView CreateInstanceOfferingEvent
type CreateInstanceOfferingEventView struct {
	Inventory InstanceOfferingInventoryView `json:"inventory,omitempty"`
}

