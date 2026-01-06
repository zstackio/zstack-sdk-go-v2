// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunRouterInterfaceInventoryView AliyunRouterInterface
type AliyunRouterInterfaceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	RouterInterfaceId string `json:"routerInterfaceId,omitempty"`
	VirtualRouterUuid string `json:"virtualRouterUuid,omitempty"`
	AccessPointUuid string `json:"accessPointUuid,omitempty"`
	Role string `json:"role,omitempty"`
	VRouterType string `json:"vRouterType,omitempty"`
	Spec string `json:"spec,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	OppositeInterfaceUuid string `json:"oppositeInterfaceUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// SyncAliyunRouterInterfaceFromRemoteEventView SyncAliyunRouterInterfaceFromRemoteEvent
type SyncAliyunRouterInterfaceFromRemoteEventView struct {
	Inventories []AliyunRouterInterfaceInventoryView `json:"inventories,omitempty"`
}

// CreateAliyunRouterInterfaceRemoteEventView CreateAliyunRouterInterfaceRemoteEvent
type CreateAliyunRouterInterfaceRemoteEventView struct {
	Inventory AliyunRouterInterfaceInventoryView `json:"inventory,omitempty"`
}

// StartConnectionBetweenAliyunRouterInterfaceEventView StartConnectionBetweenAliyunRouterInterfaceEvent
type StartConnectionBetweenAliyunRouterInterfaceEventView struct {
	Inventory AliyunRouterInterfaceInventoryView `json:"inventory,omitempty"`
}

// QueryAliyunRouterInterfaceFromLocalView QueryAliyunRouterInterfaceFromLocal
type QueryAliyunRouterInterfaceFromLocalView struct {
	Inventories []AliyunRouterInterfaceInventoryView `json:"inventories,omitempty"`
}

