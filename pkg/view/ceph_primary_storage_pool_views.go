// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CephPrimaryStoragePoolInventoryView CephPrimaryStoragePool
type CephPrimaryStoragePoolInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"poolName,omitempty"`
	rest string `json:"aliasName,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"type,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest int64 `json:"usedCapacity,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest string `json:"securityPolicy,omitempty"`
	rest int `json:"replicatedSize,omitempty"`
	rest float32 `json:"diskUtilization,omitempty"`
}

