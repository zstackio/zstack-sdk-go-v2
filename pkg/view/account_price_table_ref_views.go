// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccountPriceTableRefInventoryView AccountPriceTableRef
type AccountPriceTableRefInventoryView struct {
	AccountUuid *string `json:"accountUuid,omitempty"`
	TableUuid *string `json:"tableUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryAccountPriceTableRefView QueryAccountPriceTableRef
type QueryAccountPriceTableRefView struct {
	Inventories []AccountPriceTableRefInventoryView `json:"inventories,omitempty"`
}

// GetAccountPriceTableRefView GetAccountPriceTableRef
type GetAccountPriceTableRefView struct {
	AccountUuids []string `json:"accountUuids,omitempty"`
	TableUuid *string `json:"tableUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

