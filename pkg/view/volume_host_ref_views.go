// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeHostRefInventoryView VolumeHostRef
type VolumeHostRefInventoryView struct {
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"mountPath,omitempty"`
	rest string `json:"device,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

