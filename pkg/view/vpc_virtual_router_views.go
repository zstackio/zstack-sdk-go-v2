// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVirtualRouterInventoryView VpcVirtualRouter
type VpcVirtualRouterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VrId *string `json:"vrId,omitempty"`
	VpcUuid *string `json:"vpcUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
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

