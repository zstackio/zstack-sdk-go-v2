// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVpnConnectionInventoryView VpcVpnConnection
type VpcVpnConnectionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	ConnectionId string `json:"connectionId,omitempty"`
	UserGatewayUuid string `json:"userGatewayUuid,omitempty"`
	VpnGatewayUuid string `json:"vpnGatewayUuid,omitempty"`
	LocalSubnet string `json:"localSubnet,omitempty"`
	RemoteSubnet string `json:"remoteSubnet,omitempty"`
	IkeConfigUuid string `json:"ikeConfigUuid,omitempty"`
	IpsecConfigUuid string `json:"ipsecConfigUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

