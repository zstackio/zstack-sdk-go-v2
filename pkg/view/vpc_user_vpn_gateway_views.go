// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcUserVpnGatewayInventoryView VpcUserVpnGateway
type VpcUserVpnGatewayInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountName,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"gatewayId,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

