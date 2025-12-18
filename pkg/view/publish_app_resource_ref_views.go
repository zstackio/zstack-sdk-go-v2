// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PublishAppResourceRefInventoryView PublishAppResourceRef
type PublishAppResourceRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"appUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

