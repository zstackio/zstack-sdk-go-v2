// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SharedBlockGroupPrimaryStorageInventoryView SharedBlockGroupPrimaryStorage
type SharedBlockGroupPrimaryStorageInventoryView struct {
	rest []SharedBlockInventoryView `json:"sharedBlocks,omitempty"`
	rest string `json:"sharedBlockGroupType,omitempty"`
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

