// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcUserVpnGatewayInventoryView VpcUserVpnGateway
type VpcUserVpnGatewayInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountName *string `json:"accountName,omitempty"`
	DataCenterUuid *string `json:"dataCenterUuid,omitempty"`
	Type string `json:"type,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Description *string `json:"description,omitempty"`
}

// SyncVpcUserVpnGatewayFromRemoteEventView SyncVpcUserVpnGatewayFromRemoteEvent
type SyncVpcUserVpnGatewayFromRemoteEventView struct {
	Inventories []VpcUserVpnGatewayInventoryView `json:"inventories,omitempty"`
}

// QueryVpcUserVpnGatewayFromLocalView QueryVpcUserVpnGatewayFromLocal
type QueryVpcUserVpnGatewayFromLocalView struct {
	Inventories []VpcUserVpnGatewayInventoryView `json:"inventories,omitempty"`
}

// CreateVpcUserVpnGatewayRemoteEventView CreateVpcUserVpnGatewayRemoteEvent
type CreateVpcUserVpnGatewayRemoteEventView struct {
	Inventory VpcUserVpnGatewayInventoryView `json:"inventory,omitempty"`
}

// UpdateVpcUserVpnGatewayEventView UpdateVpcUserVpnGatewayEvent
type UpdateVpcUserVpnGatewayEventView struct {
	Inventory VpcUserVpnGatewayInventoryView `json:"inventory,omitempty"`
}

