// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VxlanHostMappingInventoryView VxlanHostMapping
type VxlanHostMappingInventoryView struct {
	rest string `json:"vxlanUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest string `json:"physicalInterface,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

