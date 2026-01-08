// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2VxlanNetworkInventoryView L2VxlanNetwork
type L2VxlanNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	Vni                  int      `json:"vni,omitempty"`
	PoolUuid             string   `json:"poolUuid,omitempty"`
	ZoneUuid             string   `json:"zoneUuid,omitempty"`
	PhysicalInterface    string   `json:"physicalInterface,omitempty"`
	Type                 string   `json:"type,omitempty"`
	VSwitchType          string   `json:"vSwitchType,omitempty"`
	VirtualNetworkId     int      `json:"virtualNetworkId,omitempty"`
	Isolated             bool     `json:"isolated,omitempty"`
	Pvlan                string   `json:"pvlan,omitempty"`
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
