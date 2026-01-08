// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ConnectionAccessPointInventoryView ConnectionAccessPoint
type ConnectionAccessPointInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccessPointId  string `json:"accessPointId,omitempty"`
	Type           string `json:"type,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	Status         string `json:"status,omitempty"`
	HostOperator   string `json:"hostOperator,omitempty"`
}

// QueryConnectionAccessPointFromLocalView QueryConnectionAccessPointFromLocal
type QueryConnectionAccessPointFromLocalView struct {
	Inventories []ConnectionAccessPointInventoryView `json:"inventories,omitempty"`
}

// GetConnectionAccessPointFromRemoteView GetConnectionAccessPointFromRemote
type GetConnectionAccessPointFromRemoteView struct {
	Inventories []ConnectionAccessPointInventoryView `json:"inventories,omitempty"`
}

// AddConnectionAccessPointFromRemoteEventView AddConnectionAccessPointFromRemoteEvent
type AddConnectionAccessPointFromRemoteEventView struct {
	Inventory ConnectionAccessPointInventoryView `json:"inventory,omitempty"`
}

// SyncConnectionAccessPointFromRemoteEventView SyncConnectionAccessPointFromRemoteEvent
type SyncConnectionAccessPointFromRemoteEventView struct {
	Inventories []ConnectionAccessPointInventoryView `json:"inventories,omitempty"`
}
