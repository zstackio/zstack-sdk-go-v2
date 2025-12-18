// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SharedBlockInventoryView SharedBlock
type SharedBlockInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"sharedBlockGroupUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"diskUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
}

