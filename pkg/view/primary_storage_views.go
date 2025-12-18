// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PrimaryStorageInventoryView PrimaryStorage
type PrimaryStorageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest int64 `json:"totalPhysicalCapacity,omitempty"`
	rest int64 `json:"availablePhysicalCapacity,omitempty"`
	rest int64 `json:"systemUsedCapacity,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"mountPath,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedClusterUuids,omitempty"`
}

