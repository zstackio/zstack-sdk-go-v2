// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VirtualBorderRouterInventoryView VirtualBorderRouter
type VirtualBorderRouterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VbrId string `json:"vbrId,omitempty"`
	VlanInterfaceId string `json:"vlanInterfaceId,omitempty"`
	Status string `json:"status,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	VlanId string `json:"vlanId,omitempty"`
	PhysicalConnectionStatus string `json:"physicalConnectionStatus,omitempty"`
	CircuitCode string `json:"circuitCode,omitempty"`
	LocalGatewayIp string `json:"localGatewayIp,omitempty"`
	PeerGatewayIp string `json:"peerGatewayIp,omitempty"`
	PeeringSubnetMask string `json:"peeringSubnetMask,omitempty"`
	PhysicalConnectionId string `json:"physicalConnectionId,omitempty"`
	AccessPointUuid string `json:"accessPointUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

