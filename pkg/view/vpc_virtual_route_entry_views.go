// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcVirtualRouteEntryInventoryView VpcVirtualRouteEntry
type VpcVirtualRouteEntryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	VRouterType string `json:"vRouterType,omitempty"`
	Status string `json:"status,omitempty"`
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty"`
	NextHopId string `json:"nextHopId,omitempty"`
	VirtualRouterUuid string `json:"virtualRouterUuid,omitempty"`
	NextHopType string `json:"nextHopType,omitempty"`
}

// CreateAliyunVpcVirtualRouterEntryRemoteEventView CreateAliyunVpcVirtualRouterEntryRemoteEvent
type CreateAliyunVpcVirtualRouterEntryRemoteEventView struct {
	Inventory VpcVirtualRouteEntryInventoryView `json:"inventory,omitempty"`
}

// QueryAliyunRouteEntryFromLocalView QueryAliyunRouteEntryFromLocal
type QueryAliyunRouteEntryFromLocalView struct {
	Inventories []VpcVirtualRouteEntryInventoryView `json:"inventories,omitempty"`
}

// SyncAliyunRouteEntryFromRemoteEventView SyncAliyunRouteEntryFromRemoteEvent
type SyncAliyunRouteEntryFromRemoteEventView struct {
	Inventories []VpcVirtualRouteEntryInventoryView `json:"inventories,omitempty"`
}

