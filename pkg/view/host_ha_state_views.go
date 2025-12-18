// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostHaStateInventoryView HostHaState
type HostHaStateInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

