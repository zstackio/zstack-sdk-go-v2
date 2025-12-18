// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VirtualBorderRouterInventoryView VirtualBorderRouter
type VirtualBorderRouterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vbrId,omitempty"`
	rest string `json:"vlanInterfaceId,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"vlanId,omitempty"`
	rest string `json:"physicalConnectionStatus,omitempty"`
	rest string `json:"circuitCode,omitempty"`
	rest string `json:"localGatewayIp,omitempty"`
	rest string `json:"peerGatewayIp,omitempty"`
	rest string `json:"peeringSubnetMask,omitempty"`
	rest string `json:"physicalConnectionId,omitempty"`
	rest string `json:"accessPointUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

