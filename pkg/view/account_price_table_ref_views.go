// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccountPriceTableRefInventoryView AccountPriceTableRef
type AccountPriceTableRefInventoryView struct {
	AccountUuid string `json:"accountUuid,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

