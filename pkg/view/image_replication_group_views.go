// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ImageReplicationGroupInventoryView ImageReplicationGroup
type ImageReplicationGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []ImageReplicationGroupBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

