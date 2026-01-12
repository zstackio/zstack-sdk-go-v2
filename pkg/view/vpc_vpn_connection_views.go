// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVpnConnectionInventoryView VpcVpnConnection
type VpcVpnConnectionInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountName *string `json:"accountName,omitempty"`
	Type string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
	ConnectionId *string `json:"connectionId,omitempty"`
	UserGatewayUuid *string `json:"userGatewayUuid,omitempty"`
	VpnGatewayUuid *string `json:"vpnGatewayUuid,omitempty"`
	LocalSubnet *string `json:"localSubnet,omitempty"`
	RemoteSubnet *string `json:"remoteSubnet,omitempty"`
	IkeConfigUuid *string `json:"ikeConfigUuid,omitempty"`
	IpsecConfigUuid *string `json:"ipsecConfigUuid,omitempty"`
}

// SyncVpcVpnConnectionFromRemoteEventView SyncVpcVpnConnectionFromRemoteEvent
type SyncVpcVpnConnectionFromRemoteEventView struct {
	Inventories []VpcVpnConnectionInventoryView `json:"inventories,omitempty"`
}

// CreateVpcVpnConnectionRemoteEventView CreateVpcVpnConnectionRemoteEvent
type CreateVpcVpnConnectionRemoteEventView struct {
	Inventory VpcVpnConnectionInventoryView `json:"inventory,omitempty"`
}

// QueryVpcVpnConnectionFromLocalView QueryVpcVpnConnectionFromLocal
type QueryVpcVpnConnectionFromLocalView struct {
	Inventories []VpcVpnConnectionInventoryView `json:"inventories,omitempty"`
}

// UpdateVpcVpnConnectionRemoteEventView UpdateVpcVpnConnectionRemoteEvent
type UpdateVpcVpnConnectionRemoteEventView struct {
	Inventory VpcVpnConnectionInventoryView `json:"inventory,omitempty"`
}

