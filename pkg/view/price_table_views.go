// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PriceTableInventoryView PriceTable
type PriceTableInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

