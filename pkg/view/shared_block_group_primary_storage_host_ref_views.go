// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SharedBlockGroupPrimaryStorageHostRefInventoryView SharedBlockGroupPrimaryStorageHostRef
type SharedBlockGroupPrimaryStorageHostRefInventoryView struct {
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest int `json:"hostId,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

