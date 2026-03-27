// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EcsImageInventoryView EcsImage
type EcsImageInventoryView struct {
	BaseInfoView
	BaseTimeView
	LocalImageUuid string `json:"localImageUuid,omitempty"`
	EcsImageId string `json:"ecsImageId,omitempty"`
	EcsImageSize int64 `json:"ecsImageSize,omitempty"`
	Description string `json:"description,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	Type string `json:"type,omitempty"`
	OssMd5Sum string `json:"ossMd5Sum,omitempty"`
	Format string `json:"format,omitempty"`
	OsName string `json:"osName,omitempty"`
}

// SyncEcsImageFromRemoteEventView SyncEcsImageFromRemoteEvent
type SyncEcsImageFromRemoteEventView struct {
	Inventories []EcsImageInventoryView `json:"inventories,omitempty"`
}

// CreateEcsImageFromEcsSnapshotEventView CreateEcsImageFromEcsSnapshotEvent
type CreateEcsImageFromEcsSnapshotEventView struct {
	Inventory EcsImageInventoryView `json:"inventory,omitempty"`
}

// CreateEcsImageFromLocalImageEventView CreateEcsImageFromLocalImageEvent
type CreateEcsImageFromLocalImageEventView struct {
	Inventory EcsImageInventoryView `json:"inventory,omitempty"`
}

// UpdateEcsImageEventView UpdateEcsImageEvent
type UpdateEcsImageEventView struct {
	Inventory EcsImageInventoryView `json:"inventory,omitempty"`
}

// QueryEcsImageFromLocalView QueryEcsImageFromLocal
type QueryEcsImageFromLocalView struct {
	Inventories []EcsImageInventoryView `json:"inventories,omitempty"`
}

