// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MiniStorageResourceReplicationInventoryView MiniStorageResourceReplication
type MiniStorageResourceReplicationInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"role,omitempty"`
	rest string `json:"networkStatus,omitempty"`
	rest string `json:"diskStatus,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

