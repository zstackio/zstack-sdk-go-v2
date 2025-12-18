// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SlbGroupL3NetworkRefInventoryView SlbGroupL3NetworkRef
type SlbGroupL3NetworkRefInventoryView struct {
	rest string `json:"slbGroupUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"l3NetworkCategory,omitempty"`
	rest string `json:"l3NetworkType,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

