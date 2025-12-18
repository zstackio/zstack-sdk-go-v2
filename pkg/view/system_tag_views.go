// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SystemTagInventoryView SystemTag
type SystemTagInventoryView struct {
	rest bool `json:"inherent,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"tag,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

