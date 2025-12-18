// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AccountPriceTableRefInventoryView AccountPriceTableRef
type AccountPriceTableRefInventoryView struct {
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"tableUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

