// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ResourceConfigInventoryView ResourceConfig
type ResourceConfigInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"category,omitempty"`
	rest string `json:"value,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

