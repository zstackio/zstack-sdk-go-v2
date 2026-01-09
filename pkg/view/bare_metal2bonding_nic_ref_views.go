// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2BondingNicRefInventoryView BareMetal2BondingNicRef
type BareMetal2BondingNicRefInventoryView struct {
	Id *int64 `json:"id,omitempty"`
	NicUuid *string `json:"nicUuid,omitempty"`
	InstanceUuid *string `json:"instanceUuid,omitempty"`
	BondingUuid *string `json:"bondingUuid,omitempty"`
	ProvisionNicUuid *string `json:"provisionNicUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	VmNic VmNicInventoryView `json:"vmNic,omitempty"`
	ProvisionNic BareMetal2InstanceProvisionNicInventoryView `json:"provisionNic,omitempty"`
	BareMetal2Bonding BareMetal2BondingInventoryView `json:"bareMetal2Bonding,omitempty"`
}

