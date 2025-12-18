// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SharedBlockCapacityInventoryView SharedBlockCapacity
type SharedBlockCapacityInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

