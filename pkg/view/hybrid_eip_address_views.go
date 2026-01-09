// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HybridEipAddressInventoryView HybridEipAddress
type HybridEipAddressInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	EipId *string `json:"eipId,omitempty"`
	BandWidth *string `json:"bandWidth,omitempty"`
	DataCenterUuid *string `json:"dataCenterUuid,omitempty"`
	AllocateResourceUuid *string `json:"allocateResourceUuid,omitempty"`
	AllocateResourceType *string `json:"allocateResourceType,omitempty"`
	Status string `json:"status,omitempty"`
	EipAddress *string `json:"eipAddress,omitempty"`
	EipType string `json:"eipType,omitempty"`
	Name string `json:"name,omitempty"`
	ChargeType *string `json:"chargeType,omitempty"`
	Description *string `json:"description,omitempty"`
	AllocateTime *time.Time `json:"allocateTime,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// SyncHybridEipFromRemoteEventView SyncHybridEipFromRemoteEvent
type SyncHybridEipFromRemoteEventView struct {
	Inventories []HybridEipAddressInventoryView `json:"inventories,omitempty"`
}

// UpdateHybridEipEventView UpdateHybridEipEvent
type UpdateHybridEipEventView struct {
	Inventory HybridEipAddressInventoryView `json:"inventory,omitempty"`
}

// QueryHybridEipFromLocalView QueryHybridEipFromLocal
type QueryHybridEipFromLocalView struct {
	Inventories []HybridEipAddressInventoryView `json:"inventories,omitempty"`
}

// AttachHybridEipToEcsEventView AttachHybridEipToEcsEvent
type AttachHybridEipToEcsEventView struct {
	Inventory HybridEipAddressInventoryView `json:"inventory,omitempty"`
}

// CreateHybridEipEventView CreateHybridEipEvent
type CreateHybridEipEventView struct {
	Inventory HybridEipAddressInventoryView `json:"inventory,omitempty"`
}

