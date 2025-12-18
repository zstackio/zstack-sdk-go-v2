// Copyright (c) ZStack.io, Inc.

package view

// GetAccountPriceTableRefView GetAccountPriceTableRef
type GetAccountPriceTableRefView struct {
	AccountUuids []string `json:"accountUuids,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

