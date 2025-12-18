// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ShareableVolumeVmInstanceRefInventoryView ShareableVolumeVmInstanceRef
type ShareableVolumeVmInstanceRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

