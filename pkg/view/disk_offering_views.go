// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DiskOfferingInventoryView DiskOffering
type DiskOfferingInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"diskSize,omitempty"`
	rest int `json:"sortKey,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"allocatorStrategy,omitempty"`
}

