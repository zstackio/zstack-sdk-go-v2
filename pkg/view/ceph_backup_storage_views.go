// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CephBackupStorageInventoryView CephBackupStorage
type CephBackupStorageInventoryView struct {
	rest []CephBackupStorageMonInventoryView `json:"mons,omitempty"`
	rest string `json:"fsid,omitempty"`
	rest string `json:"poolName,omitempty"`
	rest int64 `json:"poolAvailableCapacity,omitempty"`
	rest int64 `json:"poolUsedCapacity,omitempty"`
	rest int `json:"poolReplicatedSize,omitempty"`
	rest float32 `json:"poolDiskUtilization,omitempty"`
	rest string `json:"poolSecurityPolicy,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedZoneUuids,omitempty"`
}

