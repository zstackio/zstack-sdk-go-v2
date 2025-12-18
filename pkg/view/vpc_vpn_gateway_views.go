// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcVpnGatewayInventoryView VpcVpnGateway
type VpcVpnGatewayInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountName,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"gatewayId,omitempty"`
	rest string `json:"vSwitchUuid,omitempty"`
	rest string `json:"publicIp,omitempty"`
	rest string `json:"spec,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"businessStatus,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"endDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

