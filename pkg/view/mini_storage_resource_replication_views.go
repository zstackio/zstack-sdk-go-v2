// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MiniStorageResourceReplicationInventoryView MiniStorageResourceReplication
type MiniStorageResourceReplicationInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	Role string `json:"role,omitempty"`
	NetworkStatus string `json:"networkStatus,omitempty"`
	DiskStatus string `json:"diskStatus,omitempty"`
	Size int64 `json:"size,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

