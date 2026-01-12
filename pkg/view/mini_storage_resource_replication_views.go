// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MiniStorageResourceReplicationInventoryView MiniStorageResourceReplication
type MiniStorageResourceReplicationInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	State string `json:"state,omitempty"`
	Type *string `json:"type,omitempty"`
	Role string `json:"role,omitempty"`
	NetworkStatus string `json:"networkStatus,omitempty"`
	DiskStatus string `json:"diskStatus,omitempty"`
	Size int64 `json:"size,omitempty"`
}

// QueryMiniStorageResourceReplicationView QueryMiniStorageResourceReplication
type QueryMiniStorageResourceReplicationView struct {
	Inventories []MiniStorageResourceReplicationInventoryView `json:"inventories,omitempty"`
}

