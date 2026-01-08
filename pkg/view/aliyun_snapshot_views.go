// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunSnapshotInventoryView AliyunSnapshot
type AliyunSnapshotInventoryView struct {
	BaseInfoView
	BaseTimeView
	SnapshotId          string `json:"snapshotId,omitempty"`
	DataCenterUuid      string `json:"dataCenterUuid,omitempty"`
	DiskUuid            string `json:"diskUuid,omitempty"`
	Status              string `json:"status,omitempty"`
	AliyunSnapshotUsage string `json:"aliyunSnapshotUsage,omitempty"`
}

// UpdateAliyunSnapshotEventView UpdateAliyunSnapshotEvent
type UpdateAliyunSnapshotEventView struct {
	Inventory AliyunSnapshotInventoryView `json:"inventory,omitempty"`
}

// SyncAliyunSnapshotRemoteEventView SyncAliyunSnapshotRemoteEvent
type SyncAliyunSnapshotRemoteEventView struct {
	Inventories []AliyunSnapshotInventoryView `json:"inventories,omitempty"`
}

// QueryAliyunSnapshotFromLocalView QueryAliyunSnapshotFromLocal
type QueryAliyunSnapshotFromLocalView struct {
	Inventories []AliyunSnapshotInventoryView `json:"inventories,omitempty"`
}

// CreateAliyunSnapshotRemoteEventView CreateAliyunSnapshotRemoteEvent
type CreateAliyunSnapshotRemoteEventView struct {
	Inventory AliyunSnapshotInventoryView `json:"inventory,omitempty"`
}
