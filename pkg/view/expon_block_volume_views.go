// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ExponBlockVolumeInventoryView ExponBlockVolume
type ExponBlockVolumeInventoryView struct {
	rest string `json:"exponStatus,omitempty"`
	rest string `json:"iscsiPath,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"diskOfferingUuid,omitempty"`
	rest string `json:"rootImageUuid,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"format,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest int64 `json:"actualSize,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest bool `json:"isShareable,omitempty"`
	rest string `json:"volumeQos,omitempty"`
	rest time.Time `json:"lastDetachDate,omitempty"`
	rest string `json:"lastVmInstanceUuid,omitempty"`
	rest time.Time `json:"lastAttachDate,omitempty"`
	rest string `json:"protocol,omitempty"`
}

