// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVpnGatewayInventoryView VpcVpnGateway
type VpcVpnGatewayInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	Type string `json:"type,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
	VSwitchUuid *string `json:"vSwitchUuid,omitempty"`
	PublicIp *string `json:"publicIp,omitempty"`
	Spec *string `json:"spec,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	BusinessStatus *string `json:"businessStatus,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryVpcVpnGatewayFromLocalView QueryVpcVpnGatewayFromLocal
type QueryVpcVpnGatewayFromLocalView struct {
	Inventories []VpcVpnGatewayInventoryView `json:"inventories,omitempty"`
}

// UpdateVpcVpnGatewayEventView UpdateVpcVpnGatewayEvent
type UpdateVpcVpnGatewayEventView struct {
	Inventory VpcVpnGatewayInventoryView `json:"inventory,omitempty"`
}

// SyncVpcVpnGatewayFromRemoteEventView SyncVpcVpnGatewayFromRemoteEvent
type SyncVpcVpnGatewayFromRemoteEventView struct {
	Inventories []VpcVpnGatewayInventoryView `json:"inventories,omitempty"`
}

