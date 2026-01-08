// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VirtualBorderRouterInventoryView VirtualBorderRouter
type VirtualBorderRouterInventoryView struct {
	BaseInfoView
	BaseTimeView
	VbrId                    string `json:"vbrId,omitempty"`
	VlanInterfaceId          string `json:"vlanInterfaceId,omitempty"`
	Status                   string `json:"status,omitempty"`
	DataCenterUuid           string `json:"dataCenterUuid,omitempty"`
	VlanId                   string `json:"vlanId,omitempty"`
	PhysicalConnectionStatus string `json:"physicalConnectionStatus,omitempty"`
	CircuitCode              string `json:"circuitCode,omitempty"`
	LocalGatewayIp           string `json:"localGatewayIp,omitempty"`
	PeerGatewayIp            string `json:"peerGatewayIp,omitempty"`
	PeeringSubnetMask        string `json:"peeringSubnetMask,omitempty"`
	PhysicalConnectionId     string `json:"physicalConnectionId,omitempty"`
	AccessPointUuid          string `json:"accessPointUuid,omitempty"`
}

// QueryVirtualBorderRouterFromLocalView QueryVirtualBorderRouterFromLocal
type QueryVirtualBorderRouterFromLocalView struct {
	Inventories []VirtualBorderRouterInventoryView `json:"inventories,omitempty"`
}

// UpdateVirtualBorderRouterRemoteEventView UpdateVirtualBorderRouterRemoteEvent
type UpdateVirtualBorderRouterRemoteEventView struct {
	Inventory VirtualBorderRouterInventoryView `json:"inventory,omitempty"`
}

// SyncVirtualBorderRouterFromRemoteEventView SyncVirtualBorderRouterFromRemoteEvent
type SyncVirtualBorderRouterFromRemoteEventView struct {
	Inventories []VirtualBorderRouterInventoryView `json:"inventories,omitempty"`
}
