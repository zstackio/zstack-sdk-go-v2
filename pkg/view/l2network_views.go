// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2NetworkInventoryView L2Network
type L2NetworkInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// AttachL2NetworkToClusterEventView AttachL2NetworkToClusterEvent
type AttachL2NetworkToClusterEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// DeleteL2NetworkEventView DeleteL2NetworkEvent
type DeleteL2NetworkEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeL2NetworkVlanIdEventView ChangeL2NetworkVlanIdEvent
type ChangeL2NetworkVlanIdEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// CreateL2NetworkEventView CreateL2NetworkEvent
type CreateL2NetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// UpdateL2NetworkEventView UpdateL2NetworkEvent
type UpdateL2NetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// DetachL2NetworkFromClusterEventView DetachL2NetworkFromClusterEvent
type DetachL2NetworkFromClusterEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryL2NetworkView QueryL2Network
type QueryL2NetworkView struct {
	Inventories []L2NetworkInventoryView `json:"inventories,omitempty"`
}

