// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVirtualRouterInventoryView VpcVirtualRouter
type VpcVirtualRouterInventoryView struct {
	BaseInfoView
	BaseTimeView
	VrId    string `json:"vrId,omitempty"`
	VpcUuid string `json:"vpcUuid,omitempty"`
}

// SyncAliyunVirtualRouterFromRemoteEventView SyncAliyunVirtualRouterFromRemoteEvent
type SyncAliyunVirtualRouterFromRemoteEventView struct {
	Inventories []VpcVirtualRouterInventoryView `json:"inventories,omitempty"`
}

// UpdateAliyunVirtualRouterEventView UpdateAliyunVirtualRouterEvent
type UpdateAliyunVirtualRouterEventView struct {
	Inventory VpcVirtualRouterInventoryView `json:"inventory,omitempty"`
}

// QueryAliyunVirtualRouterFromLocalView QueryAliyunVirtualRouterFromLocal
type QueryAliyunVirtualRouterFromLocalView struct {
	Inventories []VpcVirtualRouterInventoryView `json:"inventories,omitempty"`
}
