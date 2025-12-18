// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcVpnConnectionInventoryView VpcVpnConnection
type VpcVpnConnectionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountName,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"connectionId,omitempty"`
	rest string `json:"userGatewayUuid,omitempty"`
	rest string `json:"vpnGatewayUuid,omitempty"`
	rest string `json:"localSubnet,omitempty"`
	rest string `json:"remoteSubnet,omitempty"`
	rest string `json:"ikeConfigUuid,omitempty"`
	rest string `json:"ipsecConfigUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

