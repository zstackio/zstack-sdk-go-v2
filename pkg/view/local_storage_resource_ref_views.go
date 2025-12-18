// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LocalStorageResourceRefInventoryView LocalStorageResourceRef
type LocalStorageResourceRefInventoryView struct {
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

