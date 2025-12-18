// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ZBoxInventoryView ZBox
type ZBoxInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest []ZBoxLocationRefInventoryView `json:"locationRefs,omitempty"`
	rest string `json:"mountPath,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

