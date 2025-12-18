// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ResourceDirectoryRefInventoryView ResourceDirectoryRef
type ResourceDirectoryRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"directoryUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

