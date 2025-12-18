// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVirtualRouteEntryInventoryView VpcVirtualRouteEntry
type VpcVirtualRouteEntryInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Type string `json:"type,omitempty"`
	VRouterType string `json:"vRouterType,omitempty"`
	Status string `json:"status,omitempty"`
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty"`
	NextHopId string `json:"nextHopId,omitempty"`
	VirtualRouterUuid string `json:"virtualRouterUuid,omitempty"`
	NextHopType string `json:"nextHopType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

