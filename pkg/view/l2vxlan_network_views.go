// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2VxlanNetworkInventoryView L2VxlanNetwork
type L2VxlanNetworkInventoryView struct {
	Vni *int `json:"vni,omitempty"`
	PoolUuid *string `json:"poolUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	VirtualNetworkId *int `json:"virtualNetworkId,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// CreateL2VxlanNetworkEventView CreateL2VxlanNetworkEvent
type CreateL2VxlanNetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryL2VxlanNetworkView QueryL2VxlanNetwork
type QueryL2VxlanNetworkView struct {
	Inventories []L2VxlanNetworkInventoryView `json:"inventories,omitempty"`
}

