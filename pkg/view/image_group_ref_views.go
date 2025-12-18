// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ImageGroupRefInventoryView ImageGroupRef
type ImageGroupRefInventoryView struct {
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"imageGroupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

