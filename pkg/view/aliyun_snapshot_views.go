// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunSnapshotInventoryView AliyunSnapshot
type AliyunSnapshotInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	SnapshotId string `json:"snapshotId,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	DiskUuid string `json:"diskUuid,omitempty"`
	Status string `json:"status,omitempty"`
	AliyunSnapshotUsage string `json:"aliyunSnapshotUsage,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
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

